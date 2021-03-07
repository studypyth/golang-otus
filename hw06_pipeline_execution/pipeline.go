package hw06

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// Суть функции сводится к тому, что есть канал In куда поступают входные значения для pipiline'а
// и канал Done куда может поступить Interrupt в любой момент
// Сами функции приходят в stages, мы не реализовываем логику выполнения функции stage.
// Stage принимает на вход канал со значением и отдает канал с результатами
// Мы реализуем следующий функционал. Для каждого стейджа мы запускаем функцию interStage которая запускает горутины
// Горутина делает простую вещь, перекладывает значения из канала in в канал out и прерывает если есть сообщение в канале done
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	ch := in
	for _, stage := range stages {
		ch = interStage(stage(ch), done)
	}
	return ch
}

func interStage(in In, done In) Out {
	out := make(Bi)
	go func(out Bi) {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}(out)
	return out
}
