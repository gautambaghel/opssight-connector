FROM scratch

LABEL Black Duck OpsSight Team

ARG LASTCOMMIT
ARG BUILDTIME
ARG VERSION

# Container catalog requirements
COPY ./LICENSE /licenses/
COPY ./help.1 /help.1

COPY ./opssight-quay-processor ./opssight-quay-processor

LABEL name="Black Duck OpsSight Quay Processor" \
      vendor="Black Duck Software" \
      release.version="$VERSION" \
      summary="Black Duck OpsSight Quay Processor" \
      description="Identifies image in a Quay repository upon configuring a webhook. It will send all the identified images to opssight-core for scanning. It will also retrieve vulnerability and policy violations for each image periodically from opssight-core and annotate and label each quay image with the information for each one." \
      lastcommit="$LASTCOMMIT" \
      buildtime="$BUILDTIME" \
      license="apache" \
      release="$VERSION" \
      version="$VERSION"

CMD ["./opssight-quay-processor"]
