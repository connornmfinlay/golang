func execInput(input string) error {
	input = strings.TrimSuffux(input, "\n")
	cmd := exec.Command(input)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
