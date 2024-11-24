package jsn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/exception"
	"github.com/LittleDrongo/fmn-lib/utils/files"
	"github.com/alecthomas/chroma/quick"
)

func Export(data interface{}, filepath string) error {

	files.MakeDirIfIsNotExist(filepath)

	file, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return exception.DropUp(err, "Ошибка при создании объекта данных JSON:")
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return exception.DropUp(err, "Ошибка сохранения файла JSON:")
	}

	return nil
}

func Print(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return exception.DropUp(err, "Ошибка при создании объекта данных JSON:")
	}
	fmt.Println(string(jsonData))
	return nil
}

/*
Метод печает в формате JSON соблюдая подстветку синтаксиса.

Полный список доступных стилей:

abap, algol, algol_nu, arduino, autumn, average, base16-snazzy, borland, bw, catppuccin-frappe, catppuccin-latte, catppuccin-macchiato, catppuccin-mocha, colorful, doom-one, doom-one2, dracula, emacs, evergarden, friendly, fruity, github-dark, github, gruvbox-light,  gruvbox, hr_high_contrast, hrdark, igor, lovelace, manni, modus-operandi, modus-vivendi, monokai, monokailight, murphy, native, nord, onedark, onesenterprise, paraiso-dark, paraiso-light, pastie, perldoc, pygments, rainbow_dash, rose-pine-dawn, rose-pine-moon, rose-pine, rrt, solarized-dark, solarized-dark256, solarized-light, swapoff, tango, tokyonight-day, tokyonight-moon, tokyonight-night, tokyonight-storm, trac, vim, vs, vulcan, witchhazel, xcode-dark, xcode,
*/
func ColorPrint(data interface{}, style ...string) error {

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

func ToString(data interface{}) (string, error) {
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
func Import(filepath string, anyTypePointer interface{}) error {

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
