package sftp

import (
	"fmt"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"net"
	"os"
)

/** more stuff to be used here: https://sftptogo.com/blog/go-sftp/ */

func GetConnection() *sftp.Client {

	// Get SFTP To Go URL from environment
	user := utils.Getenv("SFTP_USERNAME", "dia")
	pass := utils.Getenv("SFTP_PASSWORD", "")

	host := utils.Getenv("SFTP_URL", "ftp.indexengine.com")
	port := utils.Getenv("SFTP_PORT", "22")

	//hostKey := GetHostKey(host)

	log.Infoln("Connecting to ", host)

	var auths []ssh.AuthMethod

	if authConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(authConn).Signers))
	}

	// Use password authentication if provided
	if pass != "" {
		auths = append(auths, ssh.Password(pass))
	}

	// Initialize client configuration
	config := ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Error("Failed to connect to ", addr, err)
		os.Exit(1)
	}
	/*
		defer func(conn *ssh.Client) {
			err := conn.Close()
			if err != nil {
				log.Error(err)
			}
		}(conn)*/

	// Create new SFTP client
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Error("Unable to start SFTP subsystem: %v\n", err)
		os.Exit(1)
	}

	return sftpClient
}

/*
// GetHostKey Get host key from local known hosts
func GetHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Error("Unable to read known_hosts file: %v\n", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Error("Error parsing %q: %v\n", fields[2], err)
				os.Exit(1)
			}
			break
		}
	}

	if hostKey == nil {
		log.Error("No hostkey found for %s", host)
		os.Exit(1)
	}

	return hostKey
}
*/

func DownloadFile(sc sftp.Client, remoteFile, localFile string) (err error) {

	log.Infoln("Downloading [" + remoteFile + "] to [" + localFile + "] ")
	// Note: SFTP To Go doesn't support O_RDWR mode
	srcFile, err := sc.OpenFile(remoteFile, os.O_RDONLY)
	if err != nil {
		log.Error("Unable to open remote file: " + remoteFile)
		log.Error(err)
		return
	}
	defer func(srcFile *sftp.File) {
		err := srcFile.Close()
		if err != nil {
			log.Error(err)
		}
	}(srcFile)

	dstFile, err := os.Create(localFile)
	if err != nil {
		log.Error("Unable to open local file: " + localFile)
		log.Error(err)
		return
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			log.Error(err)
		}
	}(dstFile)

	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Error("Unable to download remote file: " + remoteFile)
		log.Error(err)
		os.Exit(1)
	}
	log.Infoln(fmt.Sprintf("%d bytes copied", bytes))

	return
}
