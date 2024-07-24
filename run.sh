#!/usr/bin/env sh
# FLOW: For each model, for each algorithm, for each file, analyze with each analysis.
SIGN='sign'
PARITY='parity'
INTERVALS='intervals'
TAINT='taint'
PENTAGONS='pentagons'
PREFIX='prefix'
SUFFIX='suffix'
TARSIS='tarsis'
analyses=($PARITY $SIGN $INTERVALS $TAINT $PENTAGONS $PREFIX $SUFFIX $TARSIS)

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


echo "Analyzing with $analysis"
for model in "${models[@]}"; do
    echo "Analyzing with $model as model."
    for algorithm in "${algorithms[@]}"; do
        echo "Analyzing with $algorithm as algorithm."
        cd "$model/$algorithm"
        file_list=($(find . -type f))
        if [ $algorithm == $ARRAY_MANIPULATION ]; then
            analyses=($PARITY $SIGN)
        elif [ $algorithm == $CUSTOM ]; then
            analyses=($PARITY $TARSIS)
        elif [ $algorithm == $SORTING ]; then
            analyses=($SUFFIX $TARSIS)
        elif [ $algorithm == $STRING_MANIPULATION ]; then
            analyses=($INTERVALS $TAINT)
        fi
        for analysis in "${analyses[@]}"; do
            for file in "${file_list[@]}"; do
                echo "Analyzing $file"
                "$GO_LISA" -i "$file" -o "$model/../output/$algorithm/$file/$analysis" -a "$analysis"
                file_count=$((file_count + 1))
            done
        done
    done
done

echo "Analyzed $file_count files."
