\# Search Term Cleaner (Go Edition)



This tool scans Google Ads accounts for poor-quality search terms, using:

\- 🔍 Hardcoded disqualifier terms (e.g. “cheap”, “diy”, “free”)

\- 🤖 OpenAI-powered AI filtering

\- 🧠 Smart account name matching

\- 🛡️ Automatic negative keyword exclusions via the Google Ads API



\## Folder Structure



```

Desktop/
└── PPC_Automation_Suite/
    └── Search_Term_Cleaner/
        ├── api/
        ├── auth_flow/
        ├── search_terms_cleaner/
        ├── shared/
        ├── go.mod
        ├── Dockerfile
        ├── .gitignore
        ├── .render
        └── README.md



\## API Endpoints



\- `GET /accounts`: Lists available account aliases

\- `GET /run-cleaner?account\_input=first doctors`: Scans the specified account



\## Deployment



Can be deployed to Render or other Go-compatible cloud services.



\## Environment Variables (set in Render or locally)



\- `GOOGLE\_ADS\_CLIENT\_ID`

\- `GOOGLE\_ADS\_CLIENT\_SECRET`

\- `GOOGLE\_ADS\_REFRESH\_TOKEN`

\- `GOOGLE\_DEVELOPER\_TOKEN`

\- `OPENAI\_API\_KEY`



\## Coming Soon



\- 🧠 Real-time AI review in GPT UI

\- ✅ Manual approval interface

\- 📥 Slack/Email summaries



