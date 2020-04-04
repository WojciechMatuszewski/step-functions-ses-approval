package emails

const (
	candidateStepOneBody = `<html>
    <body>
        <h1>Hi there {{.Email}}</h1>
        <p>We registered you within our system, please await further notifications</p>
    </body>
</html>
`
	judgeDecisionBody = `<html>
    <h1>
        Hi there
    </h1>
    <p>Candidate {{.Email}} registered</p>
    <h2>Please judge him accordingly</h2>
    <div>
        <a href="{{.AcceptEndpoint}}">To approve</a>
        <a href="{{.DenyEndpoint}}">To deny</a>
    </div>
</html>
`

	approveOutcomeBody = `
		<html>
		<h1> Hi {{.Email}} </h1>
		<p> You have been approved</p>
		</html>
`

	deniedOutcomeBody = `
		<html>
		<h1> Hi {{.Email}} </h1>
		<p> You have been denied</p>
		</html>
`
)
