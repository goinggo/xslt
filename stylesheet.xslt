<?xml version="1.0"?>
    <xsl:stylesheet version="1.0"
xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:str="http://exslt.org/strings"
extension-element-prefixes="str">
    <xsl:output method="text" />
    <xsl:template name="cleanText">
        <xsl:param name="pText" />
        <xsl:variable name="cleaned1" select="str:replace($pText, '&quot;', '')" />
<xsl:variable name="cleaned2" select="str:replace($cleaned1, '\', '')" />
<xsl:variable name="cleaned3" select="str:replace($cleaned2, '&#xa;', '')" />
<xsl:value-of select="$cleaned3" />
    </xsl:template>
    <xsl:template name="testNull">
        <xsl:param name="pText" />
        <xsl:choose>
            <xsl:when test="string-length($pText)>0">
                <xsl:value-of select="$pText"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:text>null</xsl:text>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    <xsl:template name="testNullZero">
        <xsl:param name="pText" />
        <xsl:choose>
            <xsl:when test="string-length($pText)>0">
                <xsl:value-of select="$pText"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:text>0</xsl:text>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    <xsl:template match="/">{"deals": [
    <xsl:for-each select="root/response/deals/list-item">{
        "dealid": <xsl:value-of select="id" />,
        "feed": "Yipit",
        "date_added": "<xsl:value-of select="date_added" />",
        "end_date": "<xsl:value-of select="end_date" />",
        "title": "<xsl:call-template name="cleanText"><xsl:with-param name="pText" select="title" /></xsl:call-template>",
        "description": "<xsl:call-template name="cleanText"><xsl:with-param name="pText" select="description" /></xsl:call-template>",
        "merchant_name": "<xsl:call-template name="cleanText"><xsl:with-param name="pText" select="business/name" /></xsl:call-template>",
        "url": "<xsl:value-of select="url" />",
        "image": "<xsl:value-of select="images/image_small" />",
        "price": <xsl:call-template name="testNull"><xsl:with-param name="pText" select="price/raw" /></xsl:call-template>,
        "value": <xsl:call-template name="testNull"><xsl:with-param name="pText" select="value/raw" /></xsl:call-template>,
        "discount": <xsl:call-template name="testNull"><xsl:with-param name="pText" select="discount/raw" /></xsl:call-template>,
        "source": "<xsl:call-template name="cleanText"><xsl:with-param name="pText" select="source/name" /></xsl:call-template>",
        "categories": [<xsl:for-each select="tags/list-item">"<xsl:value-of select="slug" />"<xsl:choose><xsl:when test="position() != last()">,</xsl:when></xsl:choose></xsl:for-each>],
        "locations": [<xsl:for-each select="business/locations/list-item">
            {
                "address": "<xsl:call-template name="cleanText"><xsl:with-param name="pText" select="address" /></xsl:call-template>",
                "city": "<xsl:value-of select="locality" />",
                "state": "<xsl:value-of select="state" />",
                "zip_code": "<xsl:value-of select="zip_code" />",
                "loc": [<xsl:call-template name="testNullZero"><xsl:with-param name="pText" select="lon" /></xsl:call-template>,<xsl:call-template name="testNullZero"><xsl:with-param name="pText" select="lat" /></xsl:call-template>]
            }<xsl:choose><xsl:when test="position() != last()">,</xsl:when></xsl:choose></xsl:for-each>]
        }<xsl:choose><xsl:when test="position() != last()">,
        </xsl:when></xsl:choose>
    </xsl:for-each>
]}
    </xsl:template>
</xsl:stylesheet>