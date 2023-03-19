# Resume Materi Command Line Interface (CLI)

## Command Line Interface (CLI)
Command Line cepat, powerful, pengembangan antarmuka berbasis teks digunakan untuk efektifitas, komunikasi efisien dengan computer untuk menyelesaikan serangkaian tugas yang lebih luas.

#### Mengapa menggunakan command line
<ol>
<li>Akses kontrol OS lebih mudah</li>
<li>Manajemen cepat dari jumlah yang besar dari os</li>
<li>Kemampuan untuk menympan script untuk tugas otomatis</li>
</ol>

## Beberapa Jenis Command Line Interface
<ul>
<li>Sehll = CLI for OS' Service</li>
UNIX Shell, Command Prompt (MSDOS)
<li>Other app CLI </li>
Python REPL, MYSQL CLI Client, Mongo Shell, Redis-cli dan lain-lain
</ul>

## Sehll Legend
<ul>
<li>me -> username</li>
<li>linuxbox -> hostname</li>
<li>~ -> path saat ini</li>
<li>$ -> normal user</li>
<li># -> root user</li>
</ul>

## UNIX Shell Command Yang Sering Digunakan

#### Directory
<ul>
<li>pwd</li>
<li>ls</li>
<li>mkdir</li>
<li>cd</li>
<li>rm</li>
<li>cp</li>
<li>mv</li>
<li>ln</li>
</ul>

#### Files
<ul>
<li>create : touch</li>
<li>view : cat, tail, less, head</li>
<li>editor : vim, nano</li>
<li>permission : chmod, chown</li>
<li>different : diff</li>
</ul>

#### Network
<ul>
<li>ping</li>
<li>ssh</li>
<li>netstats</li>
<li>nmap</li>
<li>ip addr</li>
<li>wget</li>
<li>curl</li>
<li>telnet</li>
<li>netcat</li>
</ul>

#### Utility
<ul>
<li>man</li>
<li>env</li>
<li>echo</li>
<li>date</li>
<li>which</li>
<li>watch</li>
<li>sudo</li>
<li>history</li>
<li>grep</li>
<li>locate</li>
</ul>

## Permission
rwx adalah tipe permission yang berlaku pada file atau folder tersebut. “rwx” awal menunjukkan hak akses untuk user. “rwx” kedua menunjukan hak akses untuk group owner, dan r-x menunjukkan hak akses untuk tipe user others.

## Shell Script
Shell, Program yang berfungsi sebagai jembatan antara pengguna dan kernel. Shell Script, bahasa pemrograman dikompilasi berdasarkan perintah shell. untuk membuat file sh dengan menggunakan perintah touch file_name.sh. Kemudian buat program menggunakan editor nano dengan mengetikkan nano nama_file.sh. Setelah di save ubah permission dengan 775 atau 755 agar file sh dapat terexecute.