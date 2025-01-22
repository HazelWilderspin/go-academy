-- /****** Microsoft SQL server Management studio 'To do' app data Drop script  ******/

USE [tempdb]
GO

IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[LIST_ITEM]') AND type in (N'U'))
DROP TABLE [dbo].LIST_ITEM

IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[LIST]') AND type in (N'U'))
DROP TABLE [dbo].LIST

IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[USER_DETAIL]') AND type in (N'U'))
DROP TABLE [dbo].USER_DETAIL

GO