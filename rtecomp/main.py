
from bs4 import BeautifulSoup
 
# Reading the data inside the xml
# file to a variable under the name
# data
with open('example/abc5/abc5.xml', 'r') as f:
    data = f.read()
 
# Passing the stored data inside
# the beautifulsoup parser, storing
# the returned object
Bs_data = BeautifulSoup(data, "xml")
 
# Finding all instances of tag
# `unique`
b_unique = Bs_data.find_all('Recover')
 

###########################################################################

import xml.etree.ElementTree as ET

# Passing the path of the
# xml document to enable the
# parsing process
tree = ET.parse('example/abc5/abc5.xml')
 
# getting the parent tag of
# the xml document
root = tree.getroot()
 
# printing the root (parent) tag
# of the xml document, along with
# its memory location
print(root)

root.tag
root.attrib
for child in root:
    print(child.tag, child.attrib)
    for childchild in child:
        print("\t", childchild.tag, childchild.attrib)

# Below is not quite working yet
# for recover in root.findall('Recover'):
#     print(recover)
