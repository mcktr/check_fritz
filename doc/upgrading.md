# Upgrading

## Upgrading to v1.1.0

### Index parameter

The parameter `-i` (`--index`) is deprecated and marked for removal in v1.2.0. It has shown that it is not reliable to 
identify smart devices by their index. Please use the new `-a` (`--ain`) parameter to specify a smart device via their 
AIN. You find the AIN in the Fritz!Box web interface in the Smart Home menu.

![AIN Number](images/upgrading-ain-number.png)

It is important to keep the space when defining the AIN parameter.

```
./check_fritz --method smart_status --ain "00000 0000000"
```