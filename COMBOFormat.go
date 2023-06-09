package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type SMTPServer struct {
	Domain string
	Host   string
	Port   int
}

type SMTPCredentials struct {
	SMTPServer SMTPServer
	Email      string
	Password   string
}

var smtpServers = map[string]SMTPServer{
    "gmail.com": {
        Domain: "gmail.com",
        Host:   "smtp.gmail.com",
        Port:   587,
    },
    "outlook.com": {
        Domain: "outlook.com",
        Host:   "smtp.office365.com",
        Port:   587,
    },
    "Live.com": {
	Domain: "live.com",
	Host: "smtp.office365.com",
	Port: 587,
    },
    "hotmail.com": {
        Domain: "hotmail.com",
        Host:   "smtp.office365.com",
	Port: 587,
    },
    "office365.com": {
        Domain: "office365.com",
        Host:   "smtp.office365.com",
        Port:   587,
    },
    "yahoo.com": {
        Domain: "yahoo.com",
        Host:   "smtp.mail.yahoo.com",
        Port:   587,
    },
    "aol.com": {
        Domain: "aol.com",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "zoho.com": {
        Domain: "zoho.com",
        Host:   "smtp.zoho.com",
        Port:   465,
    },
    "icloud.com": {
        Domain: "icloud.com",
        Host:   "smtp.mail.me.com",
        Port:   587,
    },
    "gmx.net": {
        Domain: "gmx.net",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "yandex.com": {
        Domain: "yandex.com",
        Host:   "smtp.yandex.com",
        Port:   465,
    },
    "mail.com": {
        Domain: "mail.com",
        Host:   "smtp.mail.com",
        Port:   587,
    },
    "tutanota.com": {
        Domain: "tutanota.com",
        Host:   "smtp.tutanota.com",
        Port:   587,
    },
    "ionos.com": {
        Domain: "ionos.com",
        Host:   "smtp.ionos.com",
        Port:   587,
    },
    "lycos.com": {
        Domain: "lycos.com",
        Host:   "smtp.lycos.com",
        Port:   587,
    },
    "rediffmail.com": {
        Domain: "rediffmail.com",
        Host:   "smtp.rediffmail.com",
        Port:   465,
    },
    "zimbra.com": {
        Domain: "zimbra.com",
        Host:   "smtp.zimbra.com",
        Port:   587,
    },
    "excite.com": {
        Domain: "excite.com",
        Host:   "smtp.excite.com",
        Port:   587,
    },
    "web.de": {
        Domain: "web.de",
        Host:   "smtp.web.de",
        Port:   587,
    },
    "orange.fr": {
        Domain: "orange.fr",
        Host:   "smtp.orange.fr",
        Port:   465,
    },
    "mail.ru": {
        Domain: "mail.ru",
        Host:   "smtp.mail.ru",
        Port:   465,
    },
    "t-online.de": {
        Domain: "t-online.de",
        Host:   "securemail.t-online.de",
        Port:   465,
    },
    "seznam.cz": {
        Domain: "seznam.cz",
        Host:   "smtp.seznam.cz",
        Port:   465,
    },
    "cox.net": {
        Domain: "cox.net",
        Host:   "smtp.cox.net",
        Port:   587,
    },
    "mail.ee": {
        Domain: "mail.ee",
        Host:   "smtp.mail.ee",
        Port:   465,
    },
    "mail.bg": {
        Domain: "mail.bg",
        Host:   "smtp.mail.bg",
        Port:   465,
    },
    "tiscali.it": {
        Domain: "tiscali.it",
        Host:   "smtp.tiscali.it",
        Port:   587,
    },
    "libero.it": {
        Domain: "libero.it",
        Host:   "smtp.libero.it",
        Port:   587,
    },
    "telefonica.net": {
        Domain: "telefonica.net",
        Host:   "smtp.telefonica.net",
        Port:   587,
    },
    "virgilio.it": {
        Domain: "virgilio.it",
        Host:   "smtp.virgilio.it",
        Port:   587,
    },
    "alice.it": {
        Domain: "alice.it",
        Host:   "smtp.alice.it",
        Port:   587,
    },
    "tin.it": {
        Domain: "tin.it",
        Host:   "smtp.tin.it",
        Port:   587,
    },
    "fastwebnet.it": {
        Domain: "fastwebnet.it",
        Host:   "smtp.fastwebnet.it",
        Port:   465,
    },
    "inwind.it": {
        Domain: "inwind.it",
        Host:   "smtp.inwind.it",
        Port:   587,
    },
    "iol.it": {
        Domain: "iol.it",
        Host:   "smtp.iol.it",
        Port:   587,
    },
    "email.it": {
        Domain: "email.it",
        Host:   "smtp.email.it",
        Port:   465,
    },
    "aruba.it": {
        Domain: "aruba.it",
        Host:   "smtps.aruba.it",
        Port:   465,
    },
    "telecomitalia.it": {
        Domain: "telecomitalia.it",
        Host:   "smtps.telecomitalia.it",
        Port:   465,
    },
    "vodafone.it": {
        Domain: "vodafone.it",
        Host:   "smtpmail.vodafone.it",
        Port:   465,
    },
    "hotmail.it": {
        Domain: "hotmail.it",
        Host:   "smtp.live.com",
        Port:   587,
    },
    "hotmail.co.uk": {
        Domain: "hotmail.co.uk",
        Host:   "smtp.office365.com",
        Port:   587,
    },
    "yahoo.co.in": {
        Domain: "yahoo.co.in",
        Host:   "smtp.mail.yahoo.co.in",
        Port:   587,
    },
    "yahoo.co.jp": {
        Domain: "yahoo.co.jp",
        Host:   "smtp.mail.yahoo.co.jp",
        Port:   587,
    },
    "yahoo.co.uk": {
        Domain: "yahoo.co.uk",
        Host:   "smtp.mail.yahoo.co.uk",
        Port:   587,
    },
    "yahoo.de": {
        Domain: "yahoo.de",
        Host:   "smtp.mail.yahoo.de",
        Port:   587,
    },
    "yahoo.fr": {
        Domain: "yahoo.fr",
        Host:   "smtp.mail.yahoo.fr",
        Port:   587,
    },
    "yahoo.it": {
        Domain: "yahoo.it",
        Host:   "smtp.mail.yahoo.it",
        Port:   587,
    },
    "yahoo.ca": {
        Domain: "yahoo.ca",
        Host:   "smtp.mail.yahoo.ca",
        Port:   587,
    },
    "yahoo.com.ar": {
        Domain: "yahoo.com.ar",
        Host:   "smtp.mail.yahoo.com.ar",
        Port:   587,
    },
    "yahoo.com.au": {
        Domain: "yahoo.com.au",
        Host:   "smtp.mail.yahoo.com.au",
        Port:   587,
    },
    "yahoo.com.br": {
        Domain: "yahoo.com.br",
        Host:   "smtp.mail.yahoo.com.br",
        Port:   587,
    },
    "yahoo.com.mx": {
        Domain: "yahoo.com.mx",
        Host:   "smtp.mail.yahoo.com.mx",
        Port:   587,
    },
    "yahoo.com.ph": {
        Domain: "yahoo.com.ph",
        Host:   "smtp.mail.yahoo.com.ph",
        Port:   587,
    },
    "yahoo.com.sg": {
        Domain: "yahoo.com.sg",
        Host:   "smtp.mail.yahoo.com.sg",
        Port:   587,
    },
    "yahoo.com.tw": {
        Domain: "yahoo.com.tw",
        Host:   "smtp.mail.yahoo.com.tw",
        Port:   587,
    },
    "yahoo.com.hk": {
        Domain: "yahoo.com.hk",
        Host:   "smtp.mail.yahoo.com.hk",
        Port:   587,
    },
    "yahoo.com.id": {
        Domain: "yahoo.com.id",
        Host:   "smtp.mail.yahoo.com.id",
        Port:   587,
    },
    "yahoo.com.my": {
        Domain: "yahoo.com.my",
        Host:   "smtp.mail.yahoo.com.my",
        Port:   587,
    },
    "yahoo.com.sa": {
        Domain: "yahoo.com.sa",
        Host:   "smtp.mail.yahoo.com.sa",
        Port:   587,
    },
    "yahoo.com.tr": {
        Domain: "yahoo.com.tr",
        Host:   "smtp.mail.yahoo.com.tr",
        Port:   587,
    },
    "yahoo.com.vn": {
        Domain: "yahoo.com.vn",
        Host:   "smtp.mail.yahoo.com.vn",
        Port:   587,
    },
    "aol.co.uk": {
        Domain: "aol.co.uk",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.de": {
        Domain: "aol.de",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.fr": {
        Domain: "aol.fr",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.it": {
        Domain: "aol.it",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.jp": {
        Domain: "aol.jp",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.ar": {
        Domain: "aol.com.ar",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.au": {
        Domain: "aol.com.au",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.br": {
        Domain: "aol.com.br",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.mx": {
        Domain: "aol.com.mx",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.sg": {
        Domain: "aol.com.sg",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.tw": {
        Domain: "aol.com.tw",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.hk": {
        Domain: "aol.com.hk",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.id": {
        Domain: "aol.com.id",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.my": {
        Domain: "aol.com.my",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.sa": {
        Domain: "aol.com.sa",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.tr": {
        Domain: "aol.com.tr",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "aol.com.vn": {
        Domain: "aol.com.vn",
        Host:   "smtp.aol.com",
        Port:   465,
    },
    "gmx.at": {
        Domain: "gmx.at",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.ch": {
        Domain: "gmx.ch",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.de": {
        Domain: "gmx.de",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.li": {
        Domain: "gmx.li",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.com": {
        Domain: "gmx.com",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.fr": {
        Domain: "gmx.fr",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.us": {
        Domain: "gmx.us",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.uk": {
        Domain: "gmx.co.uk",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.jp": {
        Domain: "gmx.co.jp",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.au": {
        Domain: "gmx.co.au",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.za": {
        Domain: "gmx.co.za",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.in": {
        Domain: "gmx.co.in",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.nz": {
        Domain: "gmx.co.nz",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.kr": {
        Domain: "gmx.co.kr",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.id": {
        Domain: "gmx.co.id",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.my": {
        Domain: "gmx.co.my",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.th": {
        Domain: "gmx.co.th",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    "gmx.co.vn": {
        Domain: "gmx.co.vn",
        Host:   "mail.gmx.net",
        Port:   465,
    },
    // Add more SMTP servers here
}

var mutex sync.Mutex

func resolveMXRecords(domain string) ([]*net.MX, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	return mxRecords, nil
}

func scanPort(domain string, port int) bool {
	address := fmt.Sprintf("%s:%d", domain, port)
	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func scanPorts(domain string, ports []int) (int, error) {
	for _, port := range ports {
		if scanPort(domain, port) {
			return port, nil
		}
	}

	return 0, fmt.Errorf("Failed to scan ports for: %s", domain)
}

func getDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func processEmailPassword(emailPassword string, outputFile string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	credentials := strings.Split(emailPassword, ":")
	if len(credentials) != 2 {
		log.Printf("\033[31mInvalid format: %s\033[0m\n", emailPassword)
		return
	}

	email := strings.TrimSpace(credentials[0])
	password := strings.TrimSpace(credentials[1])

	domain := getDomain(email)
	if domain == "" {
		log.Printf("\033[31mInvalid email: %s\033[0m\n", email)
		return
	}

	smtpServer, ok := smtpServers[domain]
	if !ok {
		mxRecords, err := resolveMXRecords(domain)
		if err != nil {
			log.Printf("\033[31mFailed to resolve MX records for: %s\033[0m\n", email)
			return
		}

		if len(mxRecords) == 0 {
			log.Printf("\033[31mNo MX records found for: %s\033[0m\n", email)
			return
		}

		smtpServer.Domain = domain
		smtpServer.Host = mxRecords[0].Host[:len(mxRecords[0].Host)-1] // Remove trailing dot
		smtpServer.Port = 587 // Default port for SMTP
	}

	ports := []int{587, 465, 2525, 26, 25}
	port, err := scanPorts(smtpServer.Host, ports)
	if err != nil {
		log.Printf("\033[31m%s\033[0m\n", err.Error())
		return
	}

	smtpServer.Port = port

	validFormat := isValidFormat(smtpServer, email, password)
	if validFormat {
		exists := isDuplicate(smtpServer, email, password, outputFile)
		if !exists {
			saveCredentials(smtpServer, email, password, outputFile)
			fmt.Printf("\033[32mValid: %s\033[0m\n", emailPassword)
		}
	} else {
		log.Printf("\033[31mInvalid format: %s\033[0m\n", emailPassword)
	}
}

func isDuplicate(smtpServer SMTPServer, email string, password string, outputFile string) bool {
	file, err := os.Open(outputFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 4 {
			continue
		}

		if fields[0] == smtpServer.Host && fields[1] == fmt.Sprint(smtpServer.Port) && fields[2] == email && fields[3] == password {
			return true
		}
	}

	return false
}

func isValidFormat(smtpServer SMTPServer, email string, password string) bool {
	return smtpServer.Port != 0 && email != "" && password != ""
}

func saveCredentials(smtpServer SMTPServer, email string, password string, outputFile string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Save valid SMTP credentials to the output file
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := fmt.Sprintf("%s|%d|%s|%s\n", smtpServer.Host, smtpServer.Port, email, password)
	_, err = file.WriteString(line)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	threads := flag.Int("threads", 10, "Number of concurrent threads")

	flag.Parse()

	// Check if input and output file paths are provided
	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var wg sync.WaitGroup

	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Semaphore to limit the number of concurrent threads
	sem := make(chan struct{}, *threads)

	for scanner.Scan() {
		emailPassword := scanner.Text()
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			processEmailPassword(emailPassword, *outputFile, &wg, sem)
			<-sem
		}()
	}

	wg.Wait()
}


