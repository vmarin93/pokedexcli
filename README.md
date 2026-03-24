# Pokedex CLI

A command-line Pokedex built with Go, developed as part of the [Boot.dev](https://www.boot.dev) Backend Development curriculum. This project was a key milestone in my journey to master Go, focusing on networking, concurrency, and clean project architecture.

You can follow my learning journey on my [Boot.dev profile](https://www.boot.dev/u/graciousrip58).

### Prerequisites

- [Go](https://go.dev/doc/install) 1.20 or higher installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/pokedexcli.git
   cd pokedexcli
   ```

2. Build the application:
   ```bash
   go build -o pokedex
   ```

3. Run the CLI:
   ```bash
   ./pokedex
   ```

Alternatively, you can run it directly using:
```bash
go run .
```

## Usage

Once inside the REPL, you can use the following commands:

- `help`: Displays a help message describing all available commands.
- `map`: Displays the names of 20 location areas in the Pokémon world. Each subsequent call displays the next 20 areas.
- `mapb`: Displays the previous 20 location areas.
- `explore <area_name>`: Lists all the Pokémon found in a particular area.
- `catch <pokemon_name>`: Attempts to catch a Pokémon and add it to your Pokedex.
- `inspect <pokemon_name>`: Displays the stats (height, weight, stats, types) of a Pokémon you have caught.
- `pokedex`: Lists all the Pokémon you have successfully caught.
- `exit`: Exits the Pokedex CLI.

## Technical Highlights

This project explores several core Go concepts:
- **JSON Marshaling/Unmarshaling**: Interacting with the [PokeAPI](https://pokeapi.co/).
- **HTTP Clients**: Managing network requests and timeouts.
- **Concurrency & Mutexes**: Ensuring thread safety in the `pokecache` package while handling background cleanup (reaping).
- **Project Structure**: Following Go best practices for package organization (`internal/` directory).

---
