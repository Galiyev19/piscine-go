# find . -type f -name "*.sh" | sort -r | sed "s/\.sh$//" |
find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r