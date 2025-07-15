# Example Extension

Every time [Textractor](https://github.com/Artikash/Textractor) has a sentence of text ready, it will call `ProcessSentence` on all extensions (Dlls) it finds sequentially (via `OnNewSentence`)
After the sentence has been processed by all extensions, it will be displayed.
