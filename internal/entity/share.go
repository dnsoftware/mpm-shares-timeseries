package entity

// Share Структура данных шары
type Share struct {
	UUID         string // уникальный идентификатор
	ServerID     string // идентификатор пул-сервера (типа ALEPH-1 и т.п.)
	CoinID       int64  // идентификатор монеты
	WorkerID     int64  // ID воркера
	WalletID     int64  // ID майнера (кошелька)
	ShareDate    string // время когда найдено в формате timestamp, в миллисекундах ("2006-01-02 15:04:05.999")
	Difficulty   string // сложность майнера
	Sharedif     string // сложность шары	реальная
	Nonce        string // nonce шары
	IsSolo       bool   // соло режим (оставлено для совместимости с предыдущей версией) TODO выпилить в будущем
	RewardMethod string // метод начисления вознаграждения
	Cost         string // награда за шару
}
