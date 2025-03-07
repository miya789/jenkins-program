import jenkins.model.*
import hudson.security.*

def instance = Jenkins.getInstance()
def strategy = instance.getAuthorizationStrategy()

if (strategy instanceof GlobalMatrixAuthorizationStrategy) {
	strategy.add(hudson.security.Permission.fromId("%s"), "%s")
	instance.save()
} else {
	throw new IllegalStateException("Authorization strategy is not GlobalMatrixAuthorizationStrategy")
}