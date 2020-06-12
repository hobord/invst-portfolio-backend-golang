DIR="frontend"
if [ ! -d "$DIR" ]; then
    git clone https://github.com/hobord/invst-portfolio-frontend.git frontend
    cd frontend
else
    cd frontend
    git pull
fi
yarn
yarn build
cp -r dist ../pubic
