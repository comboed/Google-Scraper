from reCaptchaV3Bypass.bypass import ReCaptchaV3Bypass

# Replace 'anchor_url' with the recaptcha anchcor URL of the page you want to bypass
target_url = "https://www.google.com/recaptcha/api2/anchor?ar=1&k=6LfwuyUTAAAAAOAmoS0fdqijC2PbbdH4kjq62Y1b&co=aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbTo0NDM.&hl=en&v=IyZ984yGrXrBd6ihLOYGwy9X&size=normal&s=xM4ZG7o4k3V8NKpTVoOohj3UAVA4-qXBOkj4HJufUAI61Krzw6v1XSy4ljCmtBwkV6lTMzENVvsHf9qABXrtF_Mb5OmHVg_MD99kQrP0z1wAT7t0mSd8uWOK4nQD15HobyDESbnjRRZCISyMpOAlWFn1pUsdpRmWzXELAnheVTPWzl0smPTdvUzK_6231OQVyYh0OditfsGNcOW5Njexi7REbpv5WUGq_q_gkdCRU-qPyzkptVRAbAPNxV9GDaRDejfvINGgInCApBiLQ8QrHxgDDvgK7pU&cb=a5bgq4fwx3tr"
bypass = ReCaptchaV3Bypass(target_url)

# Perform the bypass
gtk_value = bypass.bypass()

# Print the extracted GTK value
print(f"Extracted GTK: {gtk_value}")