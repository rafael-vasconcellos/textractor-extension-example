package main

// Exporta a função para uso externo (Go não exporta para DLL igual ao C++, mas isso seria um exemplo)
func OnNewSentence(sentencePtr *uint16, sentenceInfo *InfoForExtension) *uint16 {
	//oldSize := len(sentencePtr)
	// Processa a sentença
	if ProcessSentence(sentencePtr, NewSentenceInfo(sentenceInfo)) {
		// Se o tamanho aumentou, seria necessário realocar em C++.
		// Em Go, apenas reconverte para UTF-16.
	}

	return sentencePtr
}

func main() {}