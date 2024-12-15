<img src="https://github.com/matveynator/chicha-http-proxy/blob/main/chicha-http-proxy.png?raw=true" alt="chicha-http-proxy" width="100%" align="right" />

# **сhicha-http-proxy: Fast and Simple L3 HTTP MIRROR**

Chicha HTTP Proxy is a lightweight **Layer 3 (L3)** HTTP proxy. It forwards traffic to a target URL and supports both HTTP and HTTPS with automatic SSL certificates. Built with Go, it’s fast, easy to set up, and perfect for creating mirrors of websites or services.

---

### **Why Chicha HTTP Proxy?**
- **Simple**: Minimal setup—just specify the target URL.
- **Fast**: Written in Go for high performance.
- **Automatic SSL**: Get and renew HTTPS certificates with Let's Encrypt.
- **Cross-Platform**: Available for multiple platforms, including Linux, Windows, macOS, and more.

---

### **Download**

Choose the version for your platform:  
- **Linux**: [AMD64](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/linux/amd64/chicha-http-proxy), [ARM64](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/linux/arm64/chicha-http-proxy)  
- **Windows**: [AMD64](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/windows/amd64/chicha-http-proxy.exe), [ARM64](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/windows/arm64/chicha-http-proxy.exe)  
- **macOS**: [Intel](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/mac/amd64/chicha-http-proxy), [M1/M2](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/mac/arm64/chicha-http-proxy)

Other platforms (including FreeBSD, Solaris, and more) are available in the [full list here](http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui).

---

### **Installation**

1. **Linux/macOS**:
   ```bash
   sudo curl http://files.zabiyaka.net/chicha-http-proxy/latest/no-gui/linux/amd64/chicha-http-proxy -o /usr/local/bin/chicha-http-proxy && sudo chmod +x /usr/local/bin/chicha-http-proxy
   ```

2. **Windows**:
   Download the `.exe` file and add its directory to `PATH`.

---

### **Usage Examples**

1. **Basic HTTP Proxy**:
   Forward HTTP requests on port `8080` to a target URL:
   ```bash
   chicha-http-proxy --http-port=8080 --target-url=https://testsupport.zendesk.com
   ```

2. **HTTPS Proxy with Automatic SSL**:
   Proxy HTTPS traffic for `example.com`:
   ```bash
   chicha-http-proxy --domain=example.com --target-url=https://testsupport.zendesk.com
   ```

3. **HTTP and HTTPS Together**:
   Handle HTTP on `8080` and HTTPS on `8443`:
   ```bash
   chicha-http-proxy --http-port=8080 --https-port=8443 --domain=example.com --target-url=https://testsupport.zendesk.com
   ```

---

### **Systemd Autostart Setup**

1. **Create a Service File**:
   ```bash
   sudo mcedit /etc/systemd/system/chicha-http-proxy.service
   ```

2. **Add the Following Content**:
   ```ini
   [Unit]
   Description=Chicha HTTP Proxy
   After=network.target

   [Service]
   ExecStart=/usr/local/bin/chicha-http-proxy --http-port=8080 --domain=example.com --target-url=https://testsupport.zendesk.com
   Restart=on-failure

   [Install]
   WantedBy=multi-user.target
   ```

3. Save the file and reload systemd:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable chicha-http-proxy
   sudo systemctl start chicha-http-proxy
   ```

4. **Check the Status**:
   ```bash
   sudo systemctl status chicha-http-proxy
   ```

---

### **Key Features**
- **Automatic SSL/TLS**: Certificates issued and renewed by Let's Encrypt.
- **Customizable Ports**: Easily set HTTP (`80`) and HTTPS (`443`) ports.
- **Cross-Platform Support**: Available for a wide range of operating systems and architectures.
- **Simple Configuration**: No complicated setup—just specify your target URL.

---

Chicha HTTP Proxy is a compact and beginner-friendly tool for creating mirrors of websites or services. Download it today and get started in minutes!
