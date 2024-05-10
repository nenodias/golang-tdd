package reflexao

import "reflect"

func getValue(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Pointer {
		valor = valor.Elem()
	}
	return valor
}

func percorre(x interface{}, fn func(string)) {
	valor := getValue(x)

	quantidadeDeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		quantidadeDeValores = valor.NumField()
		obtemCampo = valor.Field
	case reflect.Slice, reflect.Array:
		quantidadeDeValores = valor.Len()
		obtemCampo = valor.Index
	case reflect.Map:
		quantidadeDeValores = len(valor.MapKeys())
		obtemCampo = func(idx int) reflect.Value {
			key := valor.MapKeys()[idx]
			return valor.MapIndex(key)
		}
	}
	for i := 0; i < quantidadeDeValores; i++ {
		percorre(obtemCampo(i).Interface(), fn)
	}
}
