# 🧾 PDF Generator API – HTML + TailwindCSS → PDF (via Browserless)

This is a Go API service that takes **raw HTML styled with TailwindCSS**, sends it to [Browserless.io](https://browserless.io/) (headless Chromium), and returns a clean PDF file.

> ✅ Best tested with TailwindCSS `v2.0.1`, but works with newer versions too.

---

## 📦 Tech Stack

- **Go** with [Fiber](https://gofiber.io/)
- **Hexagonal Architecture**
- **TailwindCSS** (you can plug your own version)
- **Browserless.io** for headless Chromium rendering

---

## 🧪 API Usage

### POST `/pdf`

**Request Body (JSON):**

```json
{
  "htmlContent": "<div class='text-2xl font-bold text-blue-500'>Hello PDF</div>"
}
```

## 🛠️ Getting Started
1. Setup .env or set env vars
```env
PORT=3000
BROWSERLESS_HOST=https://chrome.browserless.io
BROWSERLESS_TOKEN=your_browserless_api_key
TAILWINDCSS_URL=https://unpkg.com/tailwindcss@2.0.1/dist/tailwind.css
```