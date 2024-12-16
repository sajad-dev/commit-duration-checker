# Commit Duration Checker

This project is designed to check the duration of each commit in a Git repository. It helps you track how long each commit took to be completed, providing valuable insights into your project's development process.

این پروژه برای چک کردن مدت زمان هر کامیت در یک مخزن گیت هاب طراحی شده است. به شما کمک می‌کند تا مدت زمانی که هر کامیت طول کشیده است را بررسی کنید و دیدگاه‌های ارزشمندی در مورد روند توسعه پروژه‌تان بدست آورید.


## 📋 Requirements | پیش‌نیازها

To run this project, you need the following: (برای اجرای این پروژه، به موارد زیر نیاز دارید)

- Go 1.19 or higher  
- Git installed on your system  
- Supported platforms: Linux, macOS, Windows  

### 📦 Install Required Packages | نصب پکیج‌های مورد نیاز

Before running the project, you need to install the required Go packages.  

قبل از اجرای پروژه، باید پکیج‌های مورد نیاز Go را نصب کنید.

Run the following command to install the dependencies:  (برای نصب وابستگی‌ها، دستور زیر را اجرا کنید)

```bash
go mod tidy
```
## 🚀 Installation and Setup | نصب و راه‌اندازی
Follow these steps to set up the project: (برای راه‌اندازی پروژه مراحل زیر را دنبال کنید)

Clone the project: (کلون کردن پروژه)

```bash
git clone https://github.com/sajad-dev/check-status.git
```
Navigate to the project directory: (انتقال به دایرکتوری پروژه)


```bash
cd check-status
```

Install dependencies: (نصب وابستگی‌ها)

```bash
go mod tidy‍‍
```

Run the project: (اجرای پروژه)
```bash
go run main.go
```

## 📝 Available Commands | دستورات موجود
#### Here are the available commands for this project :
دستورات موجود برای این پروژه:


Run Git commands to manage your repository: (
اجرای دستورات Git برای مدیریت مخزن خود
)
```bash
duration git
 ```
Clear logs and activity : (پاک کردن لاگ‌ها و فعالیت‌ها)
```bash
duration clear 
```

Display the help message with available commands : (نمایش پیام راهنما با دستورات موجود)
```bash
duration help
 ```

Run without displaying the panel : (اجرا بدون نمایش پنل)
```bash
duration --no-pandel # Or -np
 ```

## 🧑‍💻  Author | نویسنده

Mohammad Sajad Poorajam (محمد سجاد پورعجم) 👨‍💻🚀
