GO = go
MAGE = $(GO) run github.com/magefile/mage

TARGETS := $(shell $(MAGE) -l | awk '\
	BEGIN { count = 1 }\
	/^  / { gsub(/:/, ".", $$1); if (sub(/\*$$/, "", $$1) > 0) { targets[0] = $$1 } else { targets[count++] = $$1 } }\
	END   { for (i = 0; i < count; i++) { print targets[i] } }\
')

.PHONY: $(TARGETS)

$(TARGETS):
	@$(MAGE) $(subst .,:,$@)
