package app

func (a Application) Run(args []string) error {
	return a.instance.Run(args)
}
