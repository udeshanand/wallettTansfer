

<h1> WalletTransfer</h1>

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



</body>
</html>
