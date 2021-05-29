package kv_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/kv-go"
)

func TestTextDecoder(t *testing.T) {
	suite.Run(t, &TextDecoderSuite{})
}

type TextDecoderSuite struct {
	Suite
}

type textDecoderDecodeTestCase struct {
	TestName        string
	Data            []byte
	Input           io.Reader
	Err             string
	Expected        kv.KeyValue
	ExpectedPartial []textDecoderDecodePartialCase
}

type textDecoderDecodePartialCase struct {
	Path     []string
	Type     kv.Type
	Value    string
	Children int
}

func (s *TextDecoderSuite) TestDecode() {
	testCases := []textDecoderDecodeTestCase{
		{
			TestName: "NilInput",
			Data:     nil,
			Err:      `kv: <input>:1:1: unexpected EOF`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "EmptyInput",
			Data:     []byte{},
			Err:      `kv: <input>:1:1: unexpected EOF`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "Incomplete",
			Input:    s.MustOpenFixture("sample.invalid-incomplete.txt"),
			Err:      `kv: testdata/sample.invalid-incomplete.txt:7:1: unexpected EOF`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "MissingKey",
			Input:    s.MustOpenFixture("sample.invalid-missing_key.txt"),
			Err:      `kv: testdata/sample.invalid-missing_key.txt:3:4: unexpected token "{"`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "MissingValue",
			Input:    s.MustOpenFixture("sample.invalid-missing_value.txt"),
			Err:      `kv: testdata/sample.invalid-missing_value.txt:7:4: unexpected token "}"`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "RootNonObject",
			Input:    s.MustOpenFixture("sample.invalid-non_object_root.txt"),
			Err:      `kv: testdata/sample.invalid-non_object_root.txt:2:1: unexpected EOF`,
			Expected: kv.NewKeyValueEmpty(),
		},
		{
			TestName: "Valid",
			Input:    s.MustOpenFixture("sample.valid.txt"),
			Expected: kv.NewKeyValueRoot("root").
				AddString("key", "value").
				AddString("qkey1", "value").
				AddString("qkey2", "qvalue").
				AddString("k{ey", "v}alue").
				AddString("esc_quote", `hello "world"`).
				AddString("esc.newline", "hello\nworld").
				AddString("esc,tab", "hello\tworld").
				AddString("esc*backslash", "hello\\world").
				AddString("esc", "h\\ell\to\\\"wo\n\\rld\\x01\"").
				AddString("int", "13").
				AddString("negint", "-13").
				AddString("float", "1.3").
				AddChild(
					kv.NewKeyValueObject("seq", nil).
						AddString("0", "don't!").
						AddString("1", "_second").
						AddString("2", "3"),
				).
				AddChild(
					kv.NewKeyValueObject("nonseq", nil).
						AddString("1a", "one").
						AddString("2b", "two").
						AddString("3c", "three"),
				),
		},
		// no multi-line support yet
		{
			TestName: "MultilineValue",
			Input:    s.MustOpenFixture("addon_english.txt"),
			Expected: kv.NewKeyValueEmpty(),
			Err:      `kv: error at testdata/addon_english.txt:12:96: literal not terminated`,
		},
		{
			TestName: "AddonInfo",
			Input:    s.MustOpenFixture("addoninfo.txt"),
			Expected: kv.NewKeyValueRoot("AddonInfo").
				AddChild(
					kv.NewKeyValueObject("siege02", nil).
						AddString("MaxPlayers", "5"),
				).
				AddChild(
					kv.NewKeyValueObject("hero_picker", nil).
						AddString("background_map", "scenes/darkmoon_hero_pick"),
				),
		},
		{
			TestName: "GameInfo",
			Input:    s.MustOpenFixture("gameinfo.gi"),
			ExpectedPartial: []textDecoderDecodePartialCase{
				{
					Path:     nil,
					Type:     kv.TypeObject,
					Children: 23,
				},
				{
					Path:     []string{"FileSystem"},
					Type:     kv.TypeObject,
					Children: 5,
				},
				{
					Path:     []string{"MaterialSystem2"},
					Type:     kv.TypeObject,
					Children: 1,
				},
				{
					Path:     []string{"Engine2"},
					Type:     kv.TypeObject,
					Children: 21,
				},
				{
					Path:     []string{"SceneFileCache"},
					Type:     kv.TypeObject,
					Children: 1,
				},
				{
					Path:     []string{"SceneSystem"},
					Type:     kv.TypeObject,
					Children: 6,
				},
				{
					Path:     []string{"SoundSystem"},
					Type:     kv.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"ToolsEnvironment"},
					Type:     kv.TypeObject,
					Children: 4,
				},
				{
					Path:     []string{"Hammer"},
					Type:     kv.TypeObject,
					Children: 19,
				},
				{
					Path:  []string{"Hammer", "DefaultTextureScale"},
					Type:  kv.TypeString,
					Value: "0.250000",
				},
				{
					Path:  []string{"Hammer", "TileGridBlendDefaultColor"},
					Type:  kv.TypeString,
					Value: "0 255 0",
				},
				{
					Path:     []string{"MaterialEditor"},
					Type:     kv.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"ResourceCompiler"},
					Type:     kv.TypeObject,
					Children: 2,
				},
				{
					Path:  []string{"ResourceCompiler", "DefaultMapBuilders", "gridnav"},
					Type:  kv.TypeString,
					Value: "1",
				},
				{
					Path:     []string{"RenderPipelineAliases"},
					Type:     kv.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"RenderSystem"},
					Type:     kv.TypeObject,
					Children: 1,
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.subtestDecode(testCase)
	}
}

func (s *TextDecoderSuite) subtestDecode(testCase textDecoderDecodeTestCase) {
	s.Run(testCase.TestName, func() {
		require := s.Require()
		input := testCase.Input

		if input == nil {
			input = bytes.NewReader(testCase.Data)
		}

		if closer, ok := input.(io.Closer); ok {
			defer closer.Close()
		}

		dec := kv.NewTextDecoder(input)
		actual := kv.NewKeyValueEmpty()
		err := dec.Decode(actual)

		if testCase.Err == "" {
			require.NoError(err)
		} else {
			require.EqualError(err, testCase.Err)
		}

		expected := testCase.Expected
		partial := testCase.ExpectedPartial

		if partial == nil {
			s.RequireEqualKeyValue(expected, actual)
			return
		}

		for partialCaseIdx, expectedChild := range partial {
			child := actual

			for _, key := range expectedChild.Path {
				child = child.Child(key)
				require.NotNilf(
					child,
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			}

			require.Equalf(
				expectedChild.Type,
				child.Type(),
				"child index=%d path=%q",
				partialCaseIdx,
				expectedChild.Path,
			)

			if expectedChild.Type == kv.TypeObject {
				require.Lenf(
					child.Children(),
					expectedChild.Children,
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			} else {
				require.Equalf(
					expectedChild.Value,
					child.Value(),
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			}
		}
	})
}
