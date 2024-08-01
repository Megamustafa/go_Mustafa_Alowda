#!/bin/sh

mkdir foods drinks snacks           

touch foods/menu.txt drinks/menu.txt snacks/menu.txt

echo "fried rice" "fried chicken" "chicken porridge" > foods/menu.txt
echo "coffee milk" "oat milk" "ice tea" > drinks/menu.txt

snacks=$(curl -s "https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json")
echo $snacks > snacks/menu.txt

echo "Shell Script Complete!"