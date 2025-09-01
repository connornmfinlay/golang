package utils

import "net/http"

func ParseJSON(r *http.ReadRequest(), payload any) error {
	if r.Body == nil{
		return fmt.Errorf("Missing request body")
	}

	return	err := json.NewDecoder(r.Body).Decode(payload)
}


func WriteJSON(w http.ResponseWriter, status int, v any) err {
	w.Header().Add("Content-Type", "application/json") 
	w.WriteHeader(status)
	
	return json.NewEncoder(w).Encode(v)
}



func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
