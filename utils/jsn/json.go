package jsn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/utils/files"
	"github.com/alecthomas/chroma/quick"
)

func Export(data any, filepath string) error {

	err := files.MakeDirIfIsNotExist(filepath)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return fmt.Errorf("ошибка при создании объекта данных JSON: %v", err)
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла JSON: %v", err)
	}

	return nil
}

func Print(data any) error {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return fmt.Errorf("ошибка при создании объекта данных JSON: %v", err)
	}
	fmt.Println(string(jsonData))
	return nil
}

/*
Метод печает в формате JSON соблюдая подстветку синтаксиса.

Полный список доступных стилей:

abap, algol, algol_nu, arduino, autumn, average, base16-snazzy, borland, bw, catppuccin-frappe, catppuccin-latte, catppuccin-macchiato, catppuccin-mocha, colorful, doom-one, doom-one2, dracula, emacs, evergarden, friendly, fruity, github-dark, github, gruvbox-light,  gruvbox, hr_high_contrast, hrdark, igor, lovelace, manni, modus-operandi, modus-vivendi, monokai, monokailight, murphy, native, nord, onedark, onesenterprise, paraiso-dark, paraiso-light, pastie, perldoc, pygments, rainbow_dash, rose-pine-dawn, rose-pine-moon, rose-pine, rrt, solarized-dark, solarized-dark256, solarized-light, swapoff, tango, tokyonight-day, tokyonight-moon, tokyonight-night, tokyonight-storm, trac, vim, vs, vulcan, witchhazel, xcode-dark, xcode,
*/
func ColorPrint(data any, style ...string) error {

	var theme string
	if len(style) > 0 {
		theme = style[0]
	} else {
		theme = "igor"
	}

	str, err := ToString(data)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = quick.Highlight(&buf, str, "json", "terminal", theme)
	if err != nil {
		return err
	}

	fmt.Println(buf.String())
	return nil

}

func ToString(data any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return "", fmt.Errorf("ошибка при создании объекта данных JSON: %v", err)
	}
	return string(jsonData), nil
}

/*
Сначала создаётся экземпляр класса который будет заполняться

	var myStrc myStruct
	jsn.Import("data/file.json", &myStrc)
*/
func Import(filepath string, anyTypePointer any) error {

	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &anyTypePointer)
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
