# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 0.1.4   | :white_check_mark: |
| 0.1.3   | :white_check_mark: |
| 0.1.2   | :x:                |
| < 0.1.2 | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability, please follow these steps:

### 1. Do NOT create a public GitHub issue

Security vulnerabilities should not be reported in public GitHub issues as they can be exploited before a fix is available.

### 2. Send a private report

Send an email to the maintainers with:
- Description of the vulnerability
- Steps to reproduce the issue
- Potential impact
- Any suggested fixes (if available)

### 3. Response timeline

- **Acknowledgment**: We will acknowledge your report within 48 hours
- **Initial Assessment**: We will provide an initial assessment within 5 business days
- **Regular Updates**: We will provide regular updates on progress
- **Resolution**: We aim to resolve critical vulnerabilities within 30 days

### 4. Responsible Disclosure

We follow responsible disclosure practices:
- We will work with you to understand and resolve the issue
- We will credit you for the discovery (unless you prefer to remain anonymous)
- We will coordinate the release of fixes and security advisories

## Security Best Practices

When using this SDK:

### Connection Strings
- Never hardcode connection strings in source code
- Use environment variables or secure configuration management
- Rotate keys regularly
- Use least-privilege access policies

### Network Security
- Always use HTTPS endpoints
- Implement proper certificate validation
- Consider network-level restrictions (firewalls, VPNs)

### Authentication Tokens
- Implement proper token rotation
- Use short-lived tokens when possible
- Store tokens securely (encrypted at rest)
- Implement token validation

### Input Validation
- Validate all user inputs
- Sanitize notification payloads
- Implement proper tag validation
- Use parameterized queries/requests

### Logging and Monitoring
- Avoid logging sensitive information
- Implement security monitoring
- Monitor for unusual activity patterns
- Implement proper audit trails

## Known Security Considerations

### Rate Limiting
- Implement client-side rate limiting
- Handle rate limit responses appropriately
- Consider implementing exponential backoff

### Error Handling
- Avoid exposing sensitive information in errors
- Implement proper error logging
- Handle network timeouts gracefully

### Dependencies
- Regularly update dependencies
- Monitor for security advisories
- Use dependency scanning tools

## Security Updates

Security updates will be:
- Released as soon as possible
- Communicated through GitHub Security Advisories
- Documented in the changelog
- Backwards compatible when possible

## Contact

For security-related questions or to report vulnerabilities, please contact the maintainers directly rather than using public channels. 