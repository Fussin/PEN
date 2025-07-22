package scanner

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner/modules"
	"sync"
	"time"
)

type ScanEngine struct {
	xssScanner              *modules.XSSScanner
	sqlScanner              *modules.SQLInjectionScanner
	ssrfScanner             *modules.SSRFScanner
	rceScanner              *modules.RCEScanner
	idorScanner             *modules.IDORScanner
	secretScanner           *modules.SecretExposureScanner
	xxeScanner              *modules.XXEScanner
	lfiScanner              *modules.LFIScanner
	openRedirectScanner     *modules.OpenRedirectScanner
	csrfScanner             *modules.CSRFScanner
	clickjackingScanner     *modules.ClickjackingScanner
	corsScanner             *modules.CORSScanner
	jwtScanner              *modules.JWTScanner
	graphqlScanner          *modules.GraphQLScanner
	nosqlScanner            *modules.NoSQLScanner
	ssiScanner              *modules.SSIScanner
	sstiScanner             *modules.SSTIScanner
	hostHeaderScanner       *modules.HostHeaderScanner
	httpSmugglingScanner    *modules.HTTPSmugglingScanner
	cachePoisioningScanner  *modules.CachePoisoningScanner
	websocketScanner        *modules.WebSocketScanner
	apiScanner              *modules.APIVulnScanner
	prototypeScanner        *modules.PrototypePollutionScanner
	deserializationScanner  *modules.DeserializationScanner
	xmlrpcScanner           *modules.XMLRPCScanner
	crlfiScanner            *modules.CRLFInjectionScanner
	ldapScanner             *modules.LDAPInjectionScanner
	xpathScanner            *modules.XPathInjectionScanner
	codeInjectionScanner    *modules.CodeInjectionScanner
	backupScanner           *modules.BackupFileScanner
	gitScanner              *modules.GitExposureScanner
	directoryScanner        *modules.DirectoryTraversalScanner
	configScanner           *modules.ConfigExposureScanner
}

func NewScanEngine() *ScanEngine {
	return &ScanEngine{
		xssScanner:              &modules.XSSScanner{},
		sqlScanner:              &modules.SQLInjectionScanner{},
		ssrfScanner:             &modules.SSRFScanner{},
		rceScanner:              &modules.RCEScanner{},
		idorScanner:             &modules.IDORScanner{},
		secretScanner:           &modules.SecretExposureScanner{},
		xxeScanner:              &modules.XXEScanner{},
		lfiScanner:              &modules.LFIScanner{},
		openRedirectScanner:     &modules.OpenRedirectScanner{},
		csrfScanner:             &modules.CSRFScanner{},
		clickjackingScanner:     &modules.ClickjackingScanner{},
		corsScanner:             &modules.CORSScanner{},
		jwtScanner:              &modules.JWTScanner{},
		graphqlScanner:          &modules.GraphQLScanner{},
		nosqlScanner:            &modules.NoSQLScanner{},
		ssiScanner:              &modules.SSIScanner{},
		sstiScanner:             &modules.SSTIScanner{},
		hostHeaderScanner:       &modules.HostHeaderScanner{},
		httpSmugglingScanner:    &modules.HTTPSmugglingScanner{},
		cachePoisioningScanner:  &modules.CachePoisoningScanner{},
		websocketScanner:        &modules.WebSocketScanner{},
		apiScanner:              &modules.APIVulnScanner{},
		prototypeScanner:        &modules.PrototypePollutionScanner{},
		deserializationScanner:  &modules.DeserializationScanner{},
		xmlrpcScanner:           &modules.XMLRPCScanner{},
		crlfiScanner:            &modules.CRLFInjectionScanner{},
		ldapScanner:             &modules.LDAPInjectionScanner{},
		xpathScanner:            &modules.XPathInjectionScanner{},
		codeInjectionScanner:    &modules.CodeInjectionScanner{},
		backupScanner:           &modules.BackupFileScanner{},
		gitScanner:              &modules.GitExposureScanner{},
		directoryScanner:        &modules.DirectoryTraversalScanner{},
		configScanner:           &modules.ConfigExposureScanner{},
	}
}

func (e *ScanEngine) ScanTarget(target *Target) (*ScanResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)
	defer cancel()

	result := &ScanResult{
		Target:          target,
		Vulnerabilities: make([]*Vulnerability, 0),
	}

	var wg sync.WaitGroup
	vulnChan := make(chan *Vulnerability, 1000)

	scanners := []VulnerabilityScanner{
		e.xssScanner, e.sqlScanner, e.ssrfScanner, e.rceScanner,
		e.idorScanner, e.secretScanner, e.xxeScanner, e.lfiScanner,
		e.openRedirectScanner, e.csrfScanner, e.clickjackingScanner, e.corsScanner,
		e.jwtScanner, e.graphqlScanner, e.nosqlScanner, e.ssiScanner,
		e.sstiScanner, e.hostHeaderScanner, e.httpSmugglingScanner, e.cachePoisioningScanner,
		e.websocketScanner, e.apiScanner, e.prototypeScanner, e.deserializationScanner,
		e.xmlrpcScanner, e.crlfiScanner, e.ldapScanner, e.xpathScanner,
		e.codeInjectionScanner, e.backupScanner, e.gitScanner, e.directoryScanner,
		e.configScanner,
	}

	for _, scanner := range scanners {
		wg.Add(1)
		go func(s VulnerabilityScanner) {
			defer wg.Done()
			s.Scan(ctx, target, vulnChan)
		}(scanner)
	}

	go func() {
		wg.Wait()
		close(vulnChan)
	}()

	for vuln := range vulnChan {
		result.Vulnerabilities = append(result.Vulnerabilities, vuln)
	}

	return result, nil
}
