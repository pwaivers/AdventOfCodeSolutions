def main():
    assert foo('test_input.txt') == 'fgij'

    print(foo('input.txt'))

def foo(filename):
    words = open(filename).read().splitlines()
    seen = set()
    for word in words:
        current_seen = set()
        for i in range(len(word)):
            h = hash_skip(word, i)
            if h in seen:
                return word[:i] + word[i+1:]
            current_seen.add(h)
        seen.update(current_seen)

def hash_skip(word, skip):
    mod_value = 2**31 - 1
    r_value = 31
    hash_value = 0
    for i in range(len(word)):
        if i == skip:
            continue
        hash_value = (hash_value*r_value + ord(word[i]) - ord('a')) % mod_value
    
    return hash_value


if __name__ == "__main__":
    main()