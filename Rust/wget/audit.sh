cargo run --release -- --url https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg

cargo run --release -- --url https://f-olivier.tech

cargo run --release -- --url https://golang.org/dl/go1.16.3.linux-amd64.tar.gz

cargo run --release -- --url https://assets.01-edu.org/wgetDataSamples/Sample.zip

cargo run --release -- -O=test_20MB.zip --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- -O=test_20MB.zip -P=~/Téléchargements/ --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- --rate-limit=300k --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- --rate-limit=700k --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- --rate-limit=2M --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- -i=downloads.txt

cargo run --release -- -B --url https://assets.01-edu.org/wgetDataSamples/20MB.zip

cargo run --release -- --mirror --convert-links --url http://corndog.io/

cargo run --release -- --mirror --url https://oct82.com/

cargo run --release -- --mirror --reject=gif --url https://oct82.com/

cargo run --release -- --mirror --url https://trypap.com/

cargo run --release -- --mirror -X=/img --url https://trypap.com/

cargo run --release -- --mirror --url https://theuselessweb.com/

