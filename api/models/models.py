from optparse import Option
from typing import Optional
import uuid
import datetime
from decimal import Decimal
from pydantic import BaseModel, Field, validator


class IPModel(BaseModel):
    ip_address: str = Field(index=True)
    list_name: Optional[str]  #The list this IP address was found on
    cloud_platform: Optional[str] #The platform this IP belongs to.
    ASN: Optional[str] #The BGP ASN number this address is part of.

    class Config:
        allow_population_by_field_name = True
        schema_extra = {
            "example": {
                "ip_addr":         "190.45.67.89",
                "list_name":       "https://ip-ranges.amazonaws.com/ip-ranges.json",
                "cloud_platform":  "Amazon-Web-Services"
            }
        }

