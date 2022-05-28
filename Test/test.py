def main():
    #ランダムに並べられた重複のない整数の配列
    array = [5, 4, 6, 2, 1, 9, 8, 3, 7, 10]
    # ソート実行
    sortedArray = sort(array)
    # 結果出力
    [print(i) for i in sortedArray]

def sort(array):
    # 要素が一つの場合はソートの必要がないので、そのまま返却
    if len(array) == 1:
        return array

    # 配列の先頭を基準値とする
    pivot = array[0]

    # ここから記述

    pl = 0
    pr = len(array) - 1

    while True:

        while array[pl] < pivot:
            pl += 1

        while   array[pr] > pivot:
            pr -= 1
        
        if pl < pr:
            array[pl] , array[pr] = array[pr] , array[pl]
            pl += 1
            pr -= 1
        
        else:
            break
    
    if pl == 0:
        pl += 1
    
    array[:pl] = sort(array[:pl]) 
    array[pl:] = sort(array[pl:]) 

    return array

    # ここまで記述

if __name__ == '__main__':
    main()