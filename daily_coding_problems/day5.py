def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

def car(pair):
    return pair(lambda a, b : a)


def cdr(pair):
    return pair(lambda a, b : b)


def main():
    assert 3 == car(cons(3, 4))
    assert 4 == cdr(cons(3, 4))

if __name__ == "__main__":
    main()