[![Build Status](http://qp06aslhuoa3tcqahwlpfn.us-west.webrelay.io/api/badges/NotFastEnuf/QS_Configurator/status.svg)](http://qp06aslhuoa3tcqahwlpfn.us-west.webrelay.io/NotFastEnuf/QS_Configurator)
USB configurator written by BKleiner for the QUICKSILVER Flight Controller Firmware based on Silverware by Silver13

To Build Dirty:
CREATE IMAGES IN DOCKER - from PowerShell  
```docker pull hanfer/xgo-libusb```    
```docker pull node:current```  
  
OPEN A NODE JS CONTAINER / SHELL FOR THE WEB PART OF THE CONFIG  
```docker run -it --rm -v C:\QS_Configurator:/src node:current /bin/bash```      -for windows  
```docker run -it --rm -v "$(pwd)":/src node:current /bin/bash```     -for linux current directory  
```cd /src```  
```npm --prefix web install```  
```npm --prefix web run build```  
```CTRL + D```  

BUILD THE CONFIG IN A hanfer/xgo-libusb CONTAINER / SHELL   
```docker run -it --rm -v C:\QS_Configurator:/src hanfer/xgo-libusb:latest```  
```make pkg/statik```  
```/run-env.sh windows make MODE=debug windows```   -windows app with debug - does not open the browser automatically  
```/run-env.sh windows make windows```  
```/run-env.sh darwin make darwin```  
```CTRL + D```  





