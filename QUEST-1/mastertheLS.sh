for entry in *
do
  if [ "$entry" != "." ] && [ "$entry" != ".." ]; then
    echo "$entry"
  fi
done