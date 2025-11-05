<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>WalletTransfer — README & GitHub Setup</title>
  <style>
    :root{--bg:#0f1724;--card:#0b1220;--muted:#94a3b8;--accent:#60a5fa}
    body{font-family:Inter, ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial; background:linear-gradient(180deg,#031029 0%, #071426 100%); color:#e6eef6; margin:0; padding:32px}
    .container{max-width:980px;margin:0 auto}
    header{display:flex;align-items:center;gap:16px}
    .logo{width:64px;height:64px;border-radius:12px;background:linear-gradient(135deg,var(--accent),#7dd3fc);display:flex;align-items:center;justify-content:center;font-weight:700;color:#022;box-shadow:0 8px 30px rgba(2,6,23,.6)}
    h1{margin:0;font-size:28px}
    p.lead{color:var(--muted);margin-top:6px}
    .card{background:rgba(255,255,255,0.03);padding:22px;border-radius:12px;margin-top:22px;box-shadow:0 6px 18px rgba(2,6,23,.6)}
    pre{background:rgba(0,0,0,0.35);padding:12px;border-radius:8px;overflow:auto}
    code{font-family:ui-monospace, SFMono-Regular, Menlo, Monaco, "Roboto Mono", "Courier New", monospace}
    .grid{display:grid;grid-template-columns:1fr 300px;gap:18px}
    .sidebar{background:rgba(255,255,255,0.02);padding:16px;border-radius:10px}
    .btn{display:inline-block;padding:10px 14px;border-radius:8px;background:var(--accent);color:#022;text-decoration:none;font-weight:600}
    table{width:100%;border-collapse:collapse}
    th,td{padding:8px 6px;text-align:left;border-bottom:1px solid rgba(255,255,255,0.03)}
    footer{margin-top:18px;color:var(--muted);font-size:13px}
    .muted{color:var(--muted)}
  </style>
</head>
<body>
  <div class="container">
    <header>
      <div class="logo">WT</div>
      <div>
        <h1>WalletTransfer — README & GitHub Setup</h1>
        <p class="lead">This single HTML file contains the full project README plus exact Git commands and step-by-step instructions to put the project on GitHub.</p>
      </div>
    </header>

    <div class="card">
      <h2>What you'll find inside</h2>
      <ul>
        <li>Project README (Markdown-styled content describing WalletTransfer)</li>
        <li>Complete GitHub push & setup instructions (commands to run locally)</li>
       
      </ul>

      
    </div>

    <div class="grid">
      <main class="card">

        <!-- README CONTENT START -->
        <article id="readme-markdown">
<h1>💰 WalletTransfer</h1>

<p>A simple <strong>digital wallet management system</strong> built in <strong>Go</strong> that allows users to securely <strong>add funds, transfer money between accounts, and view transaction history</strong>. This project demonstrates secure backend design, RESTful APIs, and database interaction using MySQL.</p>

<h2> Features</h2>
<ul>
  <li> <strong>User Registration & Login</strong> (session based authentication)</li>
  <li> <strong>Add Funds</strong> to wallet</li>
  <li> <strong>Transfer Funds</strong> between users</li>
  <li> <strong>Transaction History</strong> tracking</li>
  <li> <strong>Modular API design</strong> using Go packages</li>
  <li> <strong>Database connection pooling</strong> with MySQL</li>
</ul>

<h2> Tech Stack</h2>
<ul>
  <li><strong>Language:</strong> Go (Golang)</li>
  <li><strong>Frameworks/Libraries:</strong> <code>net/http</code>, <code>godotenv</code></li>
  <li><strong>Database:</strong> MySQL</li>
  <li><strong>Auth:</strong> Session based</li>
  <li><strong>Environment Config:</strong> <code>.env</code> file</li>
</ul>

<h2> Project Structure</h2>
<pre><code>wallettransfer/
│
├── helper/database.go           # Database configuration
├── handlers/                    # API handlers (business logic)
├── main.go                      # Application entry point
└── .env.example                 # Sample environment configuration
</code></pre>

<h2> Setup Instructions</h2>

<h3>1. Clone the Repository</h3>
<pre><code>git clone https://github.com/udeshanand/wallettTansfer.git
cd wallettransfer
</code></pre>

<h3>2. Create <code>.env</code> File</h3>
<pre><code>DB_USER=root
DB_PASS=yourpassword
DB_NAME=walletdb
DB_HOST=localhost
DB_PORT=3306
JWT_SECRET=yourjwtsecret
</code></pre>

<h3>3. Install Dependencies</h3>
<pre><code>go mod tidy
</code></pre>

<h3>4. Run the Server</h3>
<pre><code>go run main.go

# Expected output:
# Database connected successfully!
# Server running on port 8080
</code></pre>

<h2> API Endpoints</h2>
<table>
  <thead>
    <tr><th>Method</th><th>Endpoint</th><th>Description</th></tr>
  </thead>
  <tbody>
    <tr><td>POST</td><td><code>/register</code></td><td>Register new user</td></tr>
    <tr><td>POST</td><td><code>/login</code></td><td>Authenticate user</td></tr>
    <tr><td>POST</td><td><code>/wallet/add</code></td><td>Add funds to wallet</td></tr>
    <tr><td>POST</td><td><code>/wallet/transfer</code></td><td>Transfer funds to another user</td></tr>
    <tr><td>GET</td><td><code>/wallet/history</code></td><td>View transaction history</td></tr>
  </tbody>
</table>

<h2> Example Request (Transfer Funds)</h2>
<pre><code>POST /wallet/transfer

{
  "receiver_id": 2,
  "amount": 500
}

# Response:
{
  "message": "Transfer successful",
  "balance": 1500
}
</code></pre>

<h2> Security Highlights</h2>
<ul>
  <li>All sensitive data (passwords) stored securely via environment variables</li>
  <li>Passwords hashed with <code>bcrypt</code></li>
  <li>session-based session authentication</li>
 
</ul>

<h2> Future Enhancements</h2>
<ul>
  <li>Add admin dashboard</li>
  <li>Implement wallet-to-bank withdrawal</li>
  <li>Add email/SMS notifications</li>
  <li>Improve error handling with middleware</li>
</ul>

<h2> Author</h2>
<p><strong>Udesh Kumar</strong><br>
 <a href="mailto:udeshk28@gmail.com">Email id</a><br>
 <a href="https://github.com/udeshanand" target="_blank">GitHub Profile</a></p>

<hr>

<h2>GitHub: EXACT Commands & Steps</h2>
<p>Use these exact commands in your project folder (replace <code>yourusername</code> and other placeholders):</p>
<pre><code># 1. Initialize git (if not already):
git init

# 2. Add files:
git add .

# 3. Commit:
git commit -m "Initial commit - WalletTransfer project with README"

# 4. Create repo on GitHub (via web): https://github.com/new
#    - Name: wallettransfer
#    - Do NOT add a README there (you already have one)

# 5. Add remote and push:
git remote add origin  https://github.com/udeshanand/wallettTansfer.git
git branch -M main
git push -u origin main
</code></pre>

<h3>Editing README later</h3>
<pre><code>git add README.md
git commit -m "Update README"
git push
</code></pre>

<p class="muted">Tip: If your GitHub account uses SSH, replace the https remote URL with the SSH URL GitHub shows (ssh://...). If you run into permission errors, check that your local SSH key is added to your GitHub account or use a PAT for HTTPS pushes.</p>

        </article>
        <!-- README CONTENT END -->

      </main>

      <aside class="sidebar">
        <h3>Quick checklist</h3>
        <ol>
          <li>Create <code>README.md</code> (already in this HTML)</li>
          <li>Create <code>.env</code> from <code>.env.example</code></li>
          <li>Run <code>go mod tidy</code></li>
          <li>Run server: <code>go run main.go</code></li>
          <li>Push to GitHub using commands above</li>
        </ol>

        <h4 style="margin-top:12px">Need a README.md file too?</h4>
        <p class="muted">Use the button above to download a ready-made <code>README.md</code> copy of the content in this file.</p>
      </aside>
    </div>

    <footer>
      Created for <strong>wallettransfer</strong>. If you want a more production-ready README (badges, CI, license), tell me what CI or badges you want and I will add them.
    </footer>
  </div>


</body>
</html>
