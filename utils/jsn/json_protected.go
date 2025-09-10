package jsn

import "fmt"

func ImportProtectedStruct[S any](
	filepath string,
	tmpFactory func() any,
	apply func(tmp any) (S, error),
) (S, error) {
	var zero S

	tmp := tmpFactory()
	if tmp == nil {
		return zero, fmt.Errorf("tmpFactory is nil")
	}

	if err := Import(filepath, tmp); err != nil {
		return zero, err
	}

	out, err := apply(tmp)
	if err != nil {
		return zero, err
	}
	return out, nil
}

/*
 */
func ExportProtectedStruct[S any](
	data S,
	filepath string,
	buildTmp func(src S) (any, error),
) error {
	tmp, err := buildTmp(data)
	if err != nil {
		return err
	}
	if tmp == nil {
		return fmt.Errorf("buildTmp is nil")
	}

	return Export(tmp, filepath)
}
