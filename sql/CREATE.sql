/****** Microsoft SQL server Management studio 'To do' app data Create script  ******/

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].USER_DETAIL(
  USER_DETAIL_ID int  NOT NULL PRIMARY KEY, 
  USERNAME [nchar](32) NOT NULL UNIQUE, 
  FORENAME [nchar](32) NULL,
  SURNAME [nchar](32 ) NULL, 
  USER_PERMISSIONS [nchar](2) NULL,
  INIT_DATE DATE NULL,
) ON [PRIMARY]

CREATE TABLE [dbo].LIST(
	USER_DETAIL_ID int NOT NULL FOREIGN KEY REFERENCES [dbo].USER_DETAIL,
	LIST_ID int NOT NULL PRIMARY KEY,
	LIST_NAME nchar(30) NULL,
	LIST_INIT_DATE DATE NULL,
	LIST_IS_COMPLETE bit NOT NULL DEFAULT(0)
) ON [PRIMARY]

CREATE TABLE [dbo].LIST_ITEM(
	LIST_ID int NOT NULL FOREIGN KEY REFERENCES [dbo].LIST,
	ITEM_ID int NOT NULL PRIMARY KEY,
	ITEM_NAME			nchar(30) NULL,
	ITEM_DESC			nchar(400) NULL,
	ITEM_IS_CHECKED bit NOT NULL DEFAULT(0)
) ON [PRIMARY]

GO