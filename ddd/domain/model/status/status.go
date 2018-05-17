package status

type Status uint

const (
	Opened       Status = iota + 1 // 予約受付中
	Booked                         // 予約済み
	Confirmed                      // 予約確定
	ToBeStarted                    // 開始予定
	Started                        // 進行中
	ToBeFinished                   // 終了予定
	Finished                       // 終了
)
