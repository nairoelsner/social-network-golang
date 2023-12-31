package initialization

import (
	"os"

	"github.com/nairoelsner/socialNetworkGo/src/socialNetwork/network"
)

func Execute(network *network.Network) {
	defaultPassword := os.Getenv("DEFAULT_PASSWORD")

	network.AddUser("clarossa", defaultPassword, "Clarisse Estima")
	network.AddUser("endriys", defaultPassword, "Gabriel Endres")
	network.AddUser("n_elsner", defaultPassword, "Nairo Elsner")
	network.AddUser("seven_renato", defaultPassword, "Paulo Renato")

	network.UpdateUser("clarossa", map[string]string{"bio": ""})

	network.UpdateUser("endriys", map[string]string{"bio": ""})

	network.UpdateUser("n_elsner", map[string]string{"bio": ">hello world"})
	network.CreatePost("n_elsner", "n_elsner", "Bem-vindos ao Orbee! :)")

	network.UpdateUser("seven_renato", map[string]string{"bio": "Eu sou Paulo Renato e tenho 19 anos!"})

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

	network.AddUser("alan_turing", "senha123", "Alan Turing")
	network.AddUser("linus_torvalds", "senha123", "Linus Torvalds")
	network.AddUser("bill_gates", "senha123", "Bill Gates")
	network.AddUser("steve_jobs", "senha123", "Steve Jobs")
	network.AddUser("ada_lovelace", "senha123", "Ada Lovelace")
	network.AddUser("albert_einstein", "senha123", "Albert Einstein")
	network.AddUser("stephen_hawking", "senha123", "stephen hawking")
	network.AddUser("marie_curie", "senha123", "Marie Curie")
	network.AddUser("nikola_tesla", "senha123", "Nikola Tesla")
	network.AddUser("grace_hopper", "senha123", "Grace Hopper")
	network.AddUser("tim_berners_lee", "senha123", "Tim Berners-Lee")
	network.AddUser("margaret_hamilton", "senha123", "Margaret Hamilton")
	network.AddUser("richard_stallman", "senha123", "Richard Stallman")
	network.AddUser("neil_armstrong", "senha123", "Neil Armstrong")

	network.AddFollower("alan_turing", "linus_torvalds")
	network.AddFollower("alan_turing", "bill_gates")
	network.AddFollower("alan_turing", "steve_jobs")
	network.AddFollower("linus_torvalds", "ada_lovelace")
	network.AddFollower("bill_gates", "albert_einstein")
	network.AddFollower("bill_gates", "stephen_hawking")
	network.AddFollower("steve_jobs", "marie_curie")
	network.AddFollower("steve_jobs", "nikola_tesla")
	network.AddFollower("steve_jobs", "grace_hopper")
	network.AddFollower("ada_lovelace", "tim_berners_lee")
	network.AddFollower("albert_einstein", "margaret_hamilton")
	network.AddFollower("albert_einstein", "richard_stallman")
	network.AddFollower("albert_einstein", "neil_armstrong")
	network.AddFollower("stephen_hawking", "alan_turing")
	network.AddFollower("marie_curie", "linus_torvalds")
	network.AddFollower("marie_curie", "bill_gates")
	network.AddFollower("nikola_tesla", "steve_jobs")
	network.AddFollower("grace_hopper", "ada_lovelace")
	network.AddFollower("grace_hopper", "albert_einstein")
	network.AddFollower("tim_berners_lee", "stephen_hawking")
	network.AddFollower("tim_berners_lee", "marie_curie")
	network.AddFollower("margaret_hamilton", "nikola_tesla")
	network.AddFollower("richard_stallman", "grace_hopper")
	network.AddFollower("richard_stallman", "tim_berners_lee")
	network.AddFollower("neil_armstrong", "margaret_hamilton")
	network.AddFollower("neil_armstrong", "richard_stallman")
}
