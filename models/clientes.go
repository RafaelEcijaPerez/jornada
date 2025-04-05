package models

type Cliente struct {
    ID              string  `json:"id"`
    Name            string  `json:"name"`
    Town            string  `json:"town"`
    Address         string  `json:"address"`
    Latitud         float64 `json:"latitud"`
    Longitud        float64 `json:"longitud"`
    ArrayOptions map[string]string `json:"array_options"`
}
