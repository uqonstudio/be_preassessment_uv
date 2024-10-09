## Group Anagrams
Aplikasi ini adalah program sederhana yang digunakan untuk mengelompokkan anagram dari sebuah array string. Anagram adalah kata-kata yang memiliki huruf-huruf yang sama namun disusun dengan urutan yang berbeda.

## Cara Kerja
Program ini menggunakan pendekatan iteratif untuk mengelompokkan anagram. Pertama, program akan memeriksa setiap elemen dalam array string. Jika elemen tersebut belum dikunjungi, maka program akan membentuk sebuah kelompok anagram yang terdiri dari elemen tersebut dan elemen-elemen lain yang merupakan anagram dari elemen tersebut. Proses ini dilakukan dengan memanggil fungsi isAnagram() untuk memeriksa apakah dua string merupakan anagram satu sama lain.

## Cara Menggunakan
1. Masukkan array string yang ingin Anda kelompokkan anagram-nya ke dalam array arr pada bagian main() di dalam kode.
2. Jalankan program dengan menjalankan file Go tersebut.
3. Program akan mengelompokkan anagram dari array yang Anda masukkan dan mencetak hasilnya.
Contoh
    Input:
    ```pre
    arr := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
    ```

    Output: 
    ```pre
    [["cook" "cook"] ["save" "aves" "vase"] ["taste" "state"] ["map"]]
    ```