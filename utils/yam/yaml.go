package yam

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/utils/files"
	"github.com/alecthomas/chroma/quick"
	"gopkg.in/yaml.v2"
)

// Export writes any structure to a YAML file.
func Export(data any, filepath string) error {
	if err := files.EnsureDirForFile(filepath); err != nil {
		return err
	}

	file, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML data: %v", err)
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}

	return nil
}

// Print outputs any structure in YAML format.
func Print(data any) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML data: %v", err)
	}
	fmt.Println(string(yamlData))
	return nil
}

/*
ColorPrint outputs YAML with syntax highlighting.

Full list of available styles:

abap, algol, algol_nu, arduino, autumn, average, base16-snazzy, borland, bw, catppuccin-frappe, catppuccin-latte, catppuccin-macchiato, catppuccin-mocha, colorful, doom-one, doom-one2, dracula, emacs, evergarden, friendly, fruity, github-dark, github, gruvbox-light,  gruvbox, hr_high_contrast, hrdark, igor, lovelace, manni, modus-operandi, modus-vivendi, monokai, monokailight, murphy, native, nord, onedark, onesenterprise, paraiso-dark, paraiso-light, pastie, perldoc, pygments, rainbow_dash, rose-pine-dawn, rose-pine-moon, rose-pine, rrt, solarized-dark, solarized-dark256, solarized-light, swapoff, tango, tokyonight-day, tokyonight-moon, tokyonight-night, tokyonight-storm, trac, vim, vs, vulcan, witchhazel, xcode-dark, xcode,
*/
func ColorPrint(data any, style ...string) error {

	var theme string
	if len(style) > 0 {
		theme = style[0]
	} else {
		theme = "colorful"
	}

	str, err := ToString(data)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = quick.Highlight(&buf, str, "yaml", "terminal", theme)
	if err != nil {
		return err
	}

	fmt.Println(buf.String())
	return nil

}

func ToString(data any) (string, error) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML data: %v", err)
	}
	return string(yamlData), nil
}

/*
First create an instance of the struct that will be populated.

	var myStrc myStruct
	yam.Import("data/file.yaml", &myStrc)
*/
func Import(filepath string, pointer any) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, pointer)
	if err != nil {
		return err
	}

	return nil
}

func ImportStruct[S any](filepath string) (S, error) {
	var result S
	err := Import(filepath, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func ExportStruct[S any](data S, filepath string) error {
	err := Export(data, filepath)
	if err != nil {
		return err
	}
	return nil
}

func ImportOrCreateDefault[T any](filepath string, defaultValue T) (T, error) {
	_, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := Export(defaultValue, filepath); err != nil {
				var zero T
				return zero, err
			}
			return defaultValue, nil
		}
		var zero T
		return zero, err
	}
	return ImportStruct[T](filepath)
}
