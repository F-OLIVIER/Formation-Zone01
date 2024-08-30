cd solution && cargo build --release && cd ..

# pour l'audit
echo "\n╭───────────────────────────────────────────────╮"
echo "│                                               │"
echo "│     Robot : wall_e                            │"
echo "│     Map   : map00                             │"
echo "│                                               │"
echo "╰───────────────────────────────────────────────╯"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map00 -p1 solution/target/release/filler -p2 linux_robots/wall_e -q

echo "\n╭───────────────────────────────────────────────╮"
echo "│                                               │"
echo "│     Robot : h2_d2                             │"
echo "│     Map   : map01                             │"
echo "│                                               │"
echo "╰───────────────────────────────────────────────╯"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map01 -p1 solution/target/release/filler -p2 linux_robots/h2_d2 -q

echo "\n╭───────────────────────────────────────────────╮"
echo "│                                               │"
echo "│     Robot : bender                            │"
echo "│     Map   : map02                             │"
echo "│                                               │"
echo "╰───────────────────────────────────────────────╯"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q
echo "───────────────────────────────────────────────────────────"
./linux_game_engine -f maps/map02 -p1 solution/target/release/filler -p2 linux_robots/bender -q