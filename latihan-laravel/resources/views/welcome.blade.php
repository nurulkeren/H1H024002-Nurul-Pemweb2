<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>Profil Mahasiswa</title>

    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: Arial, sans-serif;
        }

        body {
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            background: linear-gradient(135deg, #e0f2fe, #f8fafc);
        }

        .card {
            width: 420px;
            padding: 40px;
            background: white;
            border-radius: 20px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        .icon {
            width: 80px;
            height: 80px;
            margin: 0 auto 20px;
            border-radius: 50%;
            background: #2563eb;
            color: white;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 32px;
            font-weight: bold;
        }

        h1 {
            margin-bottom: 10px;
            color: #1e293b;
            font-size: 28px;
        }

        .subtitle {
            color: #64748b;
            margin-bottom: 30px;
        }

        .info {
            text-align: left;
            background: #f8fafc;
            padding: 20px;
            border-radius: 12px;
        }

        .info-item {
            margin-bottom: 15px;
        }

        .info-item:last-child {
            margin-bottom: 0;
        }

        .label {
            display: block;
            color: #64748b;
            font-size: 14px;
            margin-bottom: 5px;
        }

        .value {
            color: #1e293b;
            font-weight: bold;
            font-size: 16px;
        }

        .footer {
            margin-top: 25px;
            color: #94a3b8;
            font-size: 13px;
        }
    </style>
</head>

<body>

    <div class="card">

        <div class="icon">
            N
        </div>

        <h1>Nurul Maftuhah</h1>

        <p class="subtitle">
            Mahasiswa Teknik Komputer
        </p>

        <div class="info">

            <div class="info-item">
                <span class="label">NIM</span>
                <span class="value">H1H024002</span>
            </div>

            <div class="info-item">
                <span class="label">Program Studi</span>
                <span class="value">Teknik Komputer</span>
            </div>

            <div class="info-item">
                <span class="label">Mata Kuliah</span>
                <span class="value">Pemrograman Web II</span>
            </div>

        </div>

        <p class="footer">
            Praktikum Modul 1 — Laravel 13
        </p>

    </div>

</body>
</html>