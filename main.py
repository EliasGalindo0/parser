from log_parser import QuakeLogParser


def main():
    parser = QuakeLogParser('qgames.txt')
    parser.get_match_summary()


if __name__ == '__main__':
    main()
