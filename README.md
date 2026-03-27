<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    
    <title>CHAT-INIT | High-Performance Go CLI Chat</title>
    <meta name="title" content="CHAT-INIT | High-Performance Go CLI Chat">
    <meta name="description" content="A real-time, WebSocket-based CLI chatting application built in Go. Optimized for Fedora Linux with a beautiful TUI using Lipgloss.">
    <meta name="keywords" content="Golang, CLI Chat, WebSockets, Fedora Linux, TUI, Lipgloss, Open Source Chat, CHAT-INIT">
    <meta name="author" content="Mikey (ui_mikey)">

    <meta property="og:type" content="website">
    <meta property="og:url" content="https://github.com/ui_mikey/CHAT-INIT">
    <meta property="og:title" content="CHAT-INIT | Pirate-Themed CLI Chat">
    <meta property="og:description" content="Real-time terminal messaging for developers. Build your crew and chat straight from the command line.">
    <meta property="og:image" content="https://raw.githubusercontent.com/ui_mikey/CHAT-INIT/main/preview.png">

    <meta property="twitter:card" content="summary_large_image">
    <meta property="twitter:title" content="CHAT-INIT | Go CLI Chat">
    <meta property="twitter:description" content="Fast, lightweight, and terminal-focused. The ultimate chat tool for Fedora users.">

    <style>
        :root {
            --bg-color: #0d1117;
            --text-color: #c9d1d9;
            --accent-purple: #8957e5;
            --accent-green: #238636;
            --terminal-black: #161b22;
        }

        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: var(--text-color);
            background-color: var(--bg-color);
            margin: 0;
            padding: 0;
        }

        .container {
            max-width: 900px;
            margin: 40px auto;
            padding: 20px;
        }

        header {
            text-align: center;
            border-bottom: 2px solid var(--accent-purple);
            padding-bottom: 20px;
            margin-bottom: 40px;
        }

        h1 {
            color: var(--accent-purple);
            font-size: 3rem;
            margin: 0;
            letter-spacing: 2px;
        }

        .badge {
            background: var(--accent-green);
            color: white;
            padding: 5px 12px;
            border-radius: 20px;
            font-size: 0.9rem;
            display: inline-block;
            margin-top: 10px;
        }

        .feature-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
            margin-top: 30px;
        }

        .card {
            background: var(--terminal-black);
            padding: 20px;
            border-radius: 8px;
            border: 1px solid #30363d;
        }

        code {
            background: #000;
            color: #58a6ff;
            padding: 2px 6px;
            border-radius: 4px;
            font-family: 'Courier New', Courier, monospace;
        }

        pre {
            background: #000;
            padding: 15px;
            border-radius: 8px;
            overflow-x: auto;
            color: #d1d5db;
            border: 1px solid var(--accent-purple);
        }

        .terminal-header {
            background: #30363d;
            padding: 5px 15px;
            border-top-left-radius: 8px;
            border-top-right-radius: 8px;
            font-size: 0.8rem;
            color: #8b949e;
        }

        footer {
            text-align: center;
            margin-top: 50px;
            font-size: 0.9rem;
            color: #8b949e;
        }

        .nav-hint {
            color: var(--accent-purple);
            font-weight: bold;
        }
    </style>
</head>
<body>

    <div class="container">
        <header>
            <h1>🏴‍☠️ CHAT-INIT</h1>
            <div class="badge">Built with Go & WebSockets</div>
            <p>A high-performance CLI chat application for the modern developer.</p>
        </header>

        <section class="card">
            <h2>🚀 Why CHAT-INIT?</h2>
            <p>Stop leaving your terminal to reply to messages. <strong>CHAT-INIT</strong> provides a lightweight, secure, and fast way to communicate with your "crew" directly from your <code>bash</code> or <code>zsh</code> shell.</p>
            <ul>
                <li><strong>Real-time Speed:</strong> Powered by Gorilla WebSockets.</li>
                <li><strong>Beautiful TUI:</strong> Styled with <code>Charmbracelet Lipgloss</code> for a premium Linux experience.</li>
                <li><strong>Privacy First:</strong> No heavy trackers—just pure Go code.</li>
            </ul>
        </section>

        <div class="feature-grid">
            <div class="card">
                <h3>💬 Direct Messaging</h3>
                <p>Simple number-based selection. Type <span class="nav-hint">0</span> to exit or pick a friend's ID to start a conversation instantly.</p>
            </div>
            <div class="card">
                <h3>⚓ Friend Management</h3>
                <p>Enter <span class="nav-hint">111</span> to manage your Pirate Crew. Add, delete, or toggle between friend lists and pending requests with ease.</p>
            </div>
        </div>

        <section style="margin-top: 40px;">
            <h2>🛠 Quick Start</h2>
            <div class="terminal-header">Fedora Terminal - Installation</div>
            <pre>
# Clone the repository
git clone https://github.com/ui_mikey/CHAT-INIT.git

# Move into the project
cd CHAT-INIT

# Launch the Client
go run client/main.go
            </pre>
        </section>

        <section class="card" style="margin-top: 40px; border-left: 4px solid var(--accent-green);">
            <h2>🏴‍☠️ Developer's Vision</h2>
            <blockquote>
                "CHAT-INIT was born from the transition of learning C++ to mastering the portability of Go. It's designed specifically for those who live in the terminal and demand efficiency." 
                <br>— <strong>Mikey (@ui_mikey)</strong>
            </blockquote>
        </section>

        <footer>
            <p>Released under the MIT License. Optimized for <strong>Fedora 43</strong>.</p>
            <p>&copy; 2026 CHAT-INIT Project by ui_mikey</p>
        </footer>
    </div>

</body>
</html>
