#!/usr/bin/env sh
SIGN='sign'
PARITY='parity'
analyses=($PARITY $SIGN)

ARRAY_MANIPULATION='/array-manipulation'
CUSTOM='/custom'
SORTING='/sorting'
STRING_MANIPULATION='/string-manipulation'
algorithms=($ARRAY_MANIPULATION $CUSTOM $SORTING $STRING_MANIPULATION)

C4AI_COMMAND_R_PLUS='/Users/visco/Documents/code/GO-Static-Analysis/c4ai-command-r-plus/src'
ZEPHYR_ORPO_141B_A35B_V01='/Users/visco/Documents/code/GO-Static-Analysis/zephyr-orpo-141b-A35b-v0.1/src'
models=($C4AI_COMMAND_R_PLUS $ZEPHYR_ORPO_141B_A35B_V01)

GO_LISA='/Users/visco/Documents/code/GO-Static-Analysis/go-lisa/go-lisa-0.1/bin/go-lisa'
current_directory=$(pwd)
file_count=0

for analysis in "${analyses[@]}"; do
    echo "Analyzing with $analysis"
    for model in "${models[@]}"; do
        echo "Analyzing with $model as model."
        for algorithm in "${algorithms[@]}"; do
            echo "Analyzing with $algorithm as algorithm."
            cd "$model/$algorithm"
            file_list=($(find . -type f))
            for file in "${file_list[@]}"; do
                echo "Analyzing $file"
                "$GO_LISA" -i "$file" -o "$model/../output/$analysis/$algorithm" -a "$analysis"
                file_count=$((file_count + 1))
            done
        done
    done
done

echo "Analyzed $file_count files."
