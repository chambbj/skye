# Skye

## What is Skye?

Skye is a command-line interface (CLI) for the [PDAL][pdal] [CLI][cli].

## Why a separate CLI?

PDAL's subcommands are just calls to it's so-called [kernels][kernel]. Many of
these kernels are mere wrappers of [filters][filter] or [pipelines][pipeline].
Unfortunately, there is a fair amount of overhead in constructing these "simple"
kernels.

## What can Skye do?

```
$ skye
skye is the main command.

Skye is a wrapper for PDAL.

Usage:
  skye [command]

Available Commands:
  colorize    Colorize point cloud
  convert     Convert point cloud formats
  crop        Crop point cloud
  drivers     Print PDAL's available drivers
  ground      Segment ground returns
  info        Report point cloud info
  pcl         Invoke PCL block
  pipeline    Pipeline
  sort        Sort point cloud
  thin        Thin point cloud
  version     Print the version number of Skye
  voxelgrid   Decimate point cloud
  help        Help about any command

Flags:
  -h, --help=false: help for skye
  -v, --view=false: View output using default application


Use "skye [command] --help" for more information about a command.
```

```
skye voxelgrid -i <input> -o <output>
```

```
skye colorize -i <input> -o <output> -r <raster>
```

which is effectively a (friendlier) wrapper for

```
pdal translate -i <input> -o <output> --filter colorization --filters.colorization.raster=<raster>
```

```
pdal info -i <input>
```

## What can Skye not do?

Skye cannot detect which filters and kernels are available at runtime. Maybe it
can someday, but for now it does not.

[cli]: https://en.wikipedia.org/wiki/Command-line_interface
[filter]: http://www.pdal.io/stages/index.html#filters
[kernel]: http://www.pdal.io/tutorial/writing-kernel.html
[pdal]: http://pdal.io
[pipeline]: http://www.pdal.io/pipeline.html
