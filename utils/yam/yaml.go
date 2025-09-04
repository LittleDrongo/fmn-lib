package yam

import (
	"bytes"
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/utils/files"
	"github.com/alecthomas/chroma/quick"
	"gopkg.in/yaml.v2"
)

// Метод экспортирует любую структуру в формате YAML файла.
func Export(data any, filepath string) error {
	files.MakeDirIfIsNotExist(filepath)

	file, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("ошибка при создании объекта данных YAML: %v", err)
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при сохранения файла YAML: %v", err)
	}

	return nil
}

// Метод печатать любую структуру в формате YAML файла.
func Print(data any) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("ошибка при создании объекта данных YAML: %v", err)
	}
	fmt.Println(string(yamlData))
	return nil
}

/*
Метод печает в формате YAML соблюдая подстветку синтаксиса.

Полный список доступных стилей:

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
		return "", fmt.Errorf("ошибка при создании объекта данных YAML: %v", err)
	}
	return string(yamlData), nil
}

/*
Сначала создаётся экземпляр класса который будет заполняться

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
