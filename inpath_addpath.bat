@inpath %1
@if errorlevel 1 goto add
@goto end

:add
@set PATH=%1;%PATH%
@goto end

:end