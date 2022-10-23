/*
1. Generar el paquete web dentro del directorio pkg.
2. Realizar la estructura Response con los capos: code, data y error:
	- code tendra el codigo de retorno.
	- data tendrá la estructura que envía la aplicación (en caso que no haya error).
	- error tendrá el error recibido en formato texto (en caso que haya error).
Desarrollar una función que reciba el code cómo entero, data como interfaz y error como string.
La función debe retornar en base al código, si es una respuesta con el data o con el error.
Implementar esta función en todos los retornos de los controladores, antes de enviar la respuesta al cliente la función debe generar la estructura que definimos.
*/
package web

type Response struct{
	Code    int `json:"Code"`
    Data    interface{} `json:"Data,omitempty"`
    Error   string `json:"Error,omitempty"` 
}

func NewResponse(code int, data interface{}, err string) Response{
	if code <300{
		return Response{
            Code: code,
            Data: data,
            Error: "",
        }
	}
	return Response{
        Code: code,
		Data: nil,
        Error: err,
	}
}