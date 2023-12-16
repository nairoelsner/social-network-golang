package initialization

import "github.com/nairoelsner/socialNetworkGo/src/socialNetwork/network"

func Execute(network *network.Network) {
	network.AddUser("clarossa", "senha123", "Clarisse Estima")
	network.AddUser("endriys", "senha123", "Gabriel Endres")
	network.AddUser("n_elsner", "senha123", "Nairo Elsner")
	network.AddUser("seven_renato", "senha123", "Paulo Renato")

	network.AddUser("bill_gates", "senha123", "Bill Gates")
	network.AddUser("steve_jobs", "senha123", "Steve Jobs")

	network.AddFollower("clarossa", "bill_gates")
	network.AddFollower("bill_gates", "steve_jobs")

	network.AddFollower("clarossa", "endriys")
	network.AddFollower("clarossa", "n_elsner")
	network.AddFollower("clarossa", "seven_renato")
	network.AddFollower("endriys", "clarossa")
	network.AddFollower("endriys", "n_elsner")
	network.AddFollower("endriys", "seven_renato")
	network.AddFollower("n_elsner", "clarossa")
	network.AddFollower("n_elsner", "endriys")
	network.AddFollower("n_elsner", "seven_renato")
	network.AddFollower("seven_renato", "clarossa")
	network.AddFollower("seven_renato", "endriys")
	network.AddFollower("seven_renato", "n_elsner")
}
